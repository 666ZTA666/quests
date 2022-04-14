package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// Ещё пара подобных задачек и Я буду просто прикреплять ссылку на hakerRank, где Я уже решал точно такую задачу.
	// Много решений Я взял оттуда, так как там подобные алгоритмические задачки Я уже решал. Поэтому комментарии не пишу.
	var v, n, s int
	for {
		fmt.Println("введите количество элементов массива")
		_, err := fmt.Scanln(&v) // считали, записали
		if err != nil {
			fmt.Println(err)
			continue
		}
		if v > 0 {
			break
		}
	}
	for {
		fmt.Println("введите максимальное значение элементов массива")
		_, err := fmt.Scanln(&n) // считали, записали
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			break
		}
	}
	for {
		fmt.Println("введите искомое значение, должно быть меньше максимального")
		_, err := fmt.Scanln(&s) // считали, записали
		if err != nil {
			fmt.Println(err)
			continue
		}
		if s >= 0 && s < n {
			break
		}
	}
	arr := make([]int, v)
	// создаем срез нужной длины
	for i := 0; i < v; i++ {
		arr[i] = rand.Intn(n)
		// заполняем "случайными" значениями
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	// сортируем через sort
	c := binarySearch(arr, s)
	// как оказалось в пакете sort есть свой бинарный поиск, тут Я оплошал и пришлось написать свой,
	// думаю хуже от этого не будет
	fmt.Println(arr)
	// выводим массив
	if c >= 0 {
		// проверка решения, если значение мы нашли, то результат будет больше или равен 0
		fmt.Println("idx =", c)
		fmt.Println("val =", arr[c])
	} else {
		// иначе мы не нашли решения.
		fmt.Println("not found")
	}
}

// sort.Search()

func binarySearch(a []int, x int) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	// начальные значения выше
	// а по ходу поиска мы их меняем
	for start <= end {
		mid := (start + end) / 2
		// берем средний индекс
		if a[mid] == x { // если нашли, выходим
			r = mid // found
			break
		} else if a[mid] < x {
			// если искомое значение больше найденной середины, то меняем начальный индекс
			start = mid + 1
		} else if a[mid] > x {
			// а иначе меняем конечный индекс
			end = mid - 1
		}
	}
	return r
}
