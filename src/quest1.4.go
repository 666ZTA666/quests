package main

import (
	"fmt"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//В данном задании Я решил, для уверенности писать все консольные вывоы еще и в файл
	w, err := os.Create("log_to_4_quest.txt") // Создаем файлик
	if err != nil {
		fmt.Println("ошибка открытия лога")
		return //проверяем на ошибку, вдруг не создался
	}
	var c int // переменная в которую мы запишем количество воркеров
	for {
		fmt.Println("введите количество воркеров)")
		_, err := fmt.Fprintln(w, "введите количество воркеров)")
		if err != nil {
			fmt.Println(err)
			return //как можно было понять не будет в этом коде непойманных ошибок
		}
		_, err = fmt.Scanln(&c)
		if err != nil {
			fmt.Println("ошибка сканера", err)
			return
		}
		_, err = fmt.Fprintln(w, c)
		if err != nil {
			fmt.Println(err)
			return
		}
		if c <= 0 {
			fmt.Println("количество воркеров должно быть больше 0")
			_, err := fmt.Fprintln(w, "количество воркеров должно быть больше 0")
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			break //Выход из колеса сансары, в случае если ввели нормальное количество воркеров
		}
	}
	mychan := make(chan int)        //Канал mychan для передачи данных в горутины
	chanelClosing := make(chan int) //канал chanelCLosing для передачи данных о завершении работы
	for i := 1; i < c+1; i++ {      // создаем С воркеров
		fmt.Print("запуск воркера №", i, "\n")
		_, err := fmt.Fprint(w, "запуск воркера №", i, "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
		go func(i int) { //каждому воркеру присвоен номер от 1 до C включительно
			var v int // чтобы в if не создавать переменные решил объявить их вне цикла
			var ok bool
			for {
				if v, ok = <-mychan; ok {
					fmt.Print("воркер №", i, ":", v, "\n")
					_, err := fmt.Fprint(w, "воркер №", i, ":", v, "\n")
					//в бесконечном цикле воркеры берут из канала число и выводят в Stdout, ну и в лог-файл
					//пока канал открыт
					if err != nil {
						fmt.Println(err)
						return
					}
				} else {
					// в случае если чтоние из канала недоступно, переменная ok == false
					fmt.Print("канал для воркера №", i, " закрыт, завершаем работу\n")
					_, err := fmt.Fprint(w, "канал для воркера №", i, " закрыт, завершаем работу\n")
					if err != nil {
						fmt.Println(err)
						return
					}
					// воркеры завершают работу, и отправляют в канал закрытия свои номера.
					chanelClosing <- i
					break //выход из очередного бесконечного цикла
				}
			}
		}(i)
	}
	//Буферизированный канал для значений сигналов закрытия
	signalChan := make(chan os.Signal, 1)
	//канал который впоследствии будут слушать в ожидании информации о закрытии всех горутин
	cleanupDone := make(chan struct{})
	//Signal.Notify отправляет в канал сигнал в случае если таковой был передан консоли
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		var i int
		for i < math.MaxInt { //в этот раз цикл не бесконечен, но тоже очень долог
			select {
			case <-signalChan: // если пришел сигнал закрытия, пишем время, закрываем канал.
				fmt.Println(time.Now())
				close(mychan)
				fmt.Println("закрываем канал передачи данных")
				_, err = fmt.Fprintln(w, "закрываем канал передачи данных")
				if err != nil {
					fmt.Println(err)
					return
				}
				// передаем информацию, что подготовительные работы по завершению программы были запущены
				cleanupDone <- struct{}{}
				return // выходим из горутины
			default:
				// Пока сигнала в канале нет, отправляем в mychan увеличивающееся значение i
				mychan <- i
				i++
			}
		}
	}()
	// если мы узнали, что подготовительные работы запущены, можем начать слушать канал,
	//в котором воркеры говорят, что уже закончили работу.
	<-cleanupDone
	//Можно было реализовать подобную механику через sync.WaitGroup, но там все горутины обезличены,
	//а если нам нужно знать о завершении работы вполне конкретных воркеров это не всегда подходит.
	for i := 0; i < c; i++ {
		val := <-chanelClosing
		fmt.Print("воркер №", val, " прекратил работу\n")
		_, err = fmt.Fprint(w, "воркер №", val, " прекратил работу\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println(time.Now())
	_, err = fmt.Fprintln(w, time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}
	err = w.Close() //закрываем лог файл.
	if err != nil {
		fmt.Println("ошибка закрытия лог файла", err)
	}
}
