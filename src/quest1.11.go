package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MyComparable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool
}

func main() {
	//Stackoverflow говорит следующее:
	//Simple Intersection: Compare each element in A to each in B (O(n^2)).
	//Hash Intersection: Put them into a hash table (O(n*x)),
	// where x is a factor of hash function efficiency (between 1 and 2).
	//Sorted Intersection: Sort A and do an optimized intersection (O(n*log(n))).
	// чуток посчитав и подумав выбираю вариант с хэшем).
	//Вариант с сортировкой работает лучше в когда (n) размер последовательностей до 7 элементов.

	var n1, n2, max int
	for {
		fmt.Println("введите количество элементов в первом слайсе")
		_, err := fmt.Scanln(&n1)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n1 > 0 {
			break
		}
	}
	for {
		fmt.Println("введите количество элементов во втором слайсе")
		_, err := fmt.Scanln(&n2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n2 > 0 {
			break
		}
	}
	for {
		fmt.Println("введите максимальное значение элемента в слайсе")
		_, err := fmt.Scanln(&max)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if max > 0 {
			break
		}
	}

	if n1 > n2 {
	} else {
		n1, n2 = n2, n1
	}
	var SliseInt1 = make([]int, n1)
	var SliseInt2 = make([]int, n2)
	for i := 0; i < n1; i++ {
		SliseInt1[i] = rand.Intn(max)
	}
	for i := 0; i < n2; i++ {
		SliseInt2[i] = rand.Intn(max) //В go нет генератора случайных строк, поэтому заполняю случайными int'ами.
	}
	Res, u, t := Hash(SliseInt1, SliseInt2)
	fmt.Println(Res)
	fmt.Println("Размер большего массива:", n1, "размер меньшего массива:", n2, "\nРазброс значений элементов массива:", max, "найдено совпадений:", u, "время работы функции:", t)
}

func Hash[T MyComparable](a []T, b []T) ([]T, int, time.Duration) {
	start := time.Now()
	set := make([]T, 0)
	hash := make(map[T]bool)
	for i := 0; i < len(a); i++ {
		hash[a[i]] = true
	}
	for i := 0; i < len(b); i++ {
		if _, found := hash[b[i]]; found {
			set = append(set, b[i])
		}
	}
	return set, len(set), time.Now().Sub(start)
}
