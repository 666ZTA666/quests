package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	fmt.Println(cap(slice))
	func(slice []string) {
		//slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
		fmt.Println(cap(slice))
	}(slice)

	fmt.Print(slice)
	fmt.Println(cap(slice))
}
