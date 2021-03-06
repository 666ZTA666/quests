package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
	//15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
	//градусов. Последовательность в подмножноствах не важна.
	myChan := make(chan float64)
	var Number int
	for {
		fmt.Println("введите количество температурных колебаний")
		_, err := fmt.Scanln(&Number) // считываем записываем
		if err != nil || Number <= 0 {
			continue
		}
		break
	}
	mapg := make(map[int][]float64)
	// создаем карту, ключ - нижняя граница интервала температур,
	// а значение массив входящих в интревал температур
	go tempGen(Number, myChan)
	//В отдельной горутине генерируем псевдослучайные значения температур, которые будут отправляться в канал
	for i := 0; i < Number; i++ {
		ValueFloat := <-myChan      // берем температуру из канала,
		ValueInt := int(ValueFloat) // отбрасываем дробную часть
		if ValueFloat < 0 {
			// если температура меньше нуля, то уменьшаем её целую часть еще на десять,
			//чтобы интервал температур от -10 до 0 был вписан в отдельный массив,
			//и не путался со значениями от 0 до 10
			ValueInt -= 10
			// дальше мы делим целую часть на 10 чтобы отбросить единицы
			ValueInt /= 10
			// и добавляем в срез мапы по ключу количества десятков новый элемент
			mapg[ValueInt*10] = append(mapg[ValueInt*10], ValueFloat)
		} else {
			// если температура больше 0, всё проще, мы делим на 10 чтобы отбросить единицы
			ValueInt /= 10
			// и добавляем в срез мапы по ключу десятков новый элемент
			mapg[ValueInt*10] = append(mapg[ValueInt*10], ValueFloat)
		}
	}
	// чтобы проверить всё ли верно получилось запускаем проверку
	for Key, ValuesFloat := range mapg {
		// в первом цикле мы проходим по карте и берем ключи и изначения, где значения это массив float
		for _, ValueFloat := range ValuesFloat {
			//во втором цикле мы проходимся по этому массиву и берем значения массива
			if ValueFloat < float64(Key) && ValueFloat >= float64(Key+10) {
				// в случае если значение массива меньше значения ключа карты
				// или если значение массива больше ключа карты более чем на 10, то печатаем ошибку с ключом
				fmt.Println("Ошибка в", Key)
				break // и заканчиваем цикл для данного массива
			}
		}
	}
	// проверка опциональная и у меня не выдает ошибок. ее можно было бы и убрать, но Я ведь старался.
	fmt.Print(mapg, "\n") // print сам сортирует мапы по ключу по возрастанию.
	// выводим мапу по порядку ключей через принт
}

// генератор температур
func tempGen(number int, myChan chan float64) {
	for i := 0; i < number; i++ { // в цикле от 0 до number
		val := float64(rand.Intn(500)) / 10 // генерируем значение от 0 до 499, превращаем во float и делим на 10
		if int(val*10)%2 == 0 {             // в случае если целая часть четная отправляем в канал как есть
			myChan <- val
		} else { // иначе меняем знак на противоположный, а именно на -, т.к. rand.Intn дает только положительные числа
			val = -val
			myChan <- val
		}
	}
	close(myChan) // в конце закрываем за собой канал
}
