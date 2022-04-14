package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//Удалить i-ый элемент из слайса.
	var x int
	for {
		fmt.Println("Введите количество элементов среза")
		_, err := fmt.Scanln(&x)
		if err != nil {
			fmt.Println(err)
		} else if x > 0 {
			break
		} else {
			fmt.Println("Число должно быть больше нуля!")
		}
	}
	arr := make([]int, x)
	for i := 0; i < x; i++ {
		arr[i] = rand.Intn(100) + 1
	}
	fmt.Println(arr)
	var y int
	for {
		fmt.Println("Введите номер(не индекс) элемента который надо удалить")
		_, err := fmt.Scanln(&y)
		if err != nil {
			fmt.Println(err)
		} else if y > 0 && y <= x {
			break
		} else if y <= 0 {
			fmt.Println("Номер != индекс, номера начинаются с 1")
		} else {
			fmt.Println("Нельзя удалить элемент, которого нет")
		}
	}
	fmt.Print("удаляем:[", arr[y-1], "]\n")
	arr1 := make([]int, x)
	copy(arr1, arr)
	fmt.Print("сохранили порядок\n", removeWithOrder(arr1, y-1), "\n")
	arr2 := make([]int, x)
	copy(arr2, arr)
	fmt.Print("нарушили порядок\n", removeWithoutOrder(arr2, y-1), "\n")

}
func removeWithOrder(s []int, i int) []int {
	v := append(s[:i], s[i+1:]...)
	return v
}
func removeWithoutOrder(s []int, i int) []int {
	v := s
	v[i] = v[len(v)-1]
	return v[:len(v)-1]
}
