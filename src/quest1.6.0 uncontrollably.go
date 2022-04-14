package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	//0. Приложение померло.
	//для реализаци первого(0) метода придется написать отдельную программу
	//не подконтрольно = не валидно
	myChan := make(chan bool)
	go func(myChan chan bool) {
		fmt.Println("Work of go-routine is beginning")
		myChan <- true
		for {
			fmt.Println("Still working")
			time.Sleep(500 * time.Millisecond)
		}
	}(myChan)
	if <-myChan {
		fmt.Println("I gonna kill you with exit")
		os.Exit(0)
	}
}
