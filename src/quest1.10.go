package main

import (
	"fmt"
	"math"
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
		_, err := fmt.Scanln(&Number)
		if err != nil || Number <= 0 {
			continue
		}
		break
	}
	mapg := make(map[int][]float64)
	go tempGen(Number, myChan)
	for i := 0; i < Number; i++ {
		ValueFloat := <-myChan
		ValueInt := int(ValueFloat)
		if ValueFloat < 0 {
			ValueInt -= 10
			ValueInt /= 10
			mapg[ValueInt*10] = append(mapg[ValueInt*10], ValueFloat)
		} else {
			ValueInt /= 10
			mapg[ValueInt*10] = append(mapg[ValueInt*10], ValueFloat)
		}
	}
	for Key, ValuesFloat := range mapg {
		for _, ValueFloat := range ValuesFloat {
			if ValueFloat < float64(Key) && ValueFloat >= float64(Key+10) {
				fmt.Println("Ошибка в", Key)
				break
			}
		}
	}
	fmt.Print(mapg, "\n") // print сам сортирует мапы по ключу по возрастанию.
}

func tempGen(number int, myChan chan float64) {
	for i := 0; i < number; i++ {
		val := float64(rand.Intn(500)) / 10
		if int(math.Trunc(val*10))%2 == 0 {
			myChan <- val
		} else {
			val = -val
			myChan <- val
		}
	}
	close(myChan)
}
