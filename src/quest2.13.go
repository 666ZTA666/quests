package main

import "fmt"

//todo почему int 8 при append увеличивает capacity до 10, int 8 до 12, а unt 8 до 8
func someAction(v []int8, b int8) {
	v[0] = 100
	fmt.Println(cap(v))
	v = append(v, b)
	fmt.Println(cap(v))
}
func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}
