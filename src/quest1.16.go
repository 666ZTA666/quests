package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var v, n int
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
	arr := make([]int, v)
	for i := 0; i < v; i++ {
		arr[i] = rand.Intn(n)
	}
	fmt.Println(arr)
	start := time.Now()
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr)
	fmt.Println(time.Now().Sub(start))
}

/* "своя" реализация, переписал sort.Slice под int
func quickSort(arr []int) []int {
	qsort(arr, 0, len(arr)-1)

	return arr
}

func qsort(arr []int, lowIdx, highIdx int) {
	if lowIdx < highIdx {
		pivotUnit := partition(arr, lowIdx, highIdx)
		qsort(arr, lowIdx, pivotUnit-1)
		qsort(arr, pivotUnit+1, highIdx)
	}
}

func partition(arr []int, lowIdx, highIdx int) int {
	pivotUnit := arr[highIdx]
	i := lowIdx - 1

	for j := lowIdx; j < highIdx; j++ {
		if arr[j] <= pivotUnit {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[highIdx] = arr[highIdx], arr[i+1]

	return i + 1
}
*/
