package main

import (
	"fmt"
	"sync"
	"time"
)

func Sqad(u int) {
	fmt.Println(u * u)
}
func main() {
	fmt.Println(time.Now())
	var i = make([]int, 5, 5)
	wg := new(sync.WaitGroup)
	for j := 1; j <= len(i); j++ {
		i[j-1] = j * 2
		//fmt.Println(i[j-1])
	}
	wg.Add(5)
	for j := 0; j < 5; j++ {
		go func(j int, i []int) {
			Sqad(i[j])
			wg.Done()
		}(j, i)
	}
	wg.Wait()
	fmt.Println(time.Now())
}
