package main

import (
	"fmt"
)

func main() {
	mychan := make(chan int)
	myxchan := make(chan int)
	exitchan := make(chan bool)
	go func(mychan chan int) {
		for i := 0; i <= 100; i++ {
			mychan <- i
			//time.Sleep(100 * time.Millisecond)
		}
		fmt.Println("First worker is over. First chanel closed.")
		close(mychan)
	}(mychan)
	go func(mychan, myxchan chan int) {
		for val := range mychan {
			myxchan <- val * val
		}
		fmt.Println("Second worker is over. Second chanel closed.")
		close(myxchan)
	}(mychan, myxchan)

	go func(myxchan chan int, exitchan chan bool) {
		for val := range myxchan {
			fmt.Println(val)
		}
		fmt.Println("All work is over.")
		exitchan <- true
	}(myxchan, exitchan)
	<-exitchan
}
