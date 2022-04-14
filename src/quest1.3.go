package main

import (
	"fmt"
	"sync"
	"time"
)

func Sqads(u int) (v int) {
	v = u * u
	fmt.Println(v)
	return v
}
func main() {
	fmt.Println(time.Now())
	var i = make([]int, 5, 5)
	wg := new(sync.WaitGroup)
	for j := 1; j <= len(i); j++ {
		i[j-1] = j * 2
	}
	fmt.Println(time.Now())
	wg.Add(5)
	var v int
	for j := 0; j < 5; j++ {
		go func(j int, i []int, v *int) {
			*v += Sqads(i[j])
			wg.Done()
		}(j, i, &v)
	}
	wg.Wait()
	fmt.Println(v)
	fmt.Println(time.Now())
}
