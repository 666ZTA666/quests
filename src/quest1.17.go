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
		_, err := fmt.Scanln(&v)
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
		_, err := fmt.Scanln(&n)
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
		_, err := fmt.Scanln(&s)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if s >= 0 && s < n {
			break
		}
	}
	arr := make([]int, v)
	for i := 0; i < v; i++ {
		arr[i] = rand.Intn(n)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	c := binarySearch(arr, s)
	fmt.Println(arr)
	if c >= 0 {
		fmt.Println("idx =", c)
		fmt.Println("val =", arr[c])
	} else {
		fmt.Println("not found")
	}
}

// sort.Search()

func binarySearch(a []int, x int) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			r = mid // found
			break
		} else if a[mid] < x {
			start = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}
	return r
}
