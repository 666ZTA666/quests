package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x int64
	for {
		fmt.Println("введите первое число")
		_, err := fmt.Scanln(&x) // считали, записали 1
		if err != nil {
			fmt.Println(err)
		}
		if x != 0 {
			break
		} else {
			fmt.Println("Делить на ноль мы не будем") // проверка на 0
		}
	}
	var y int64
	for {
		fmt.Println("введите второе число")
		_, err := fmt.Scanln(&y) // считали, записали 2
		if err != nil {
			fmt.Println(err)
		}
		if y != 0 {
			break
		} else {
			fmt.Println("Делить на ноль мы не будем") // проверка на 0
		}
	}

	X := big.NewInt(x) // создали бигинт 1
	Y := big.NewInt(y) // создали бигинт 2
	fmt.Println("1 =", X, "2 =", Y)
	// вывели их

	Z := big.NewInt(0).Add(Y, X)
	fmt.Println("1+2 =", Z) // сумма
	Z = Z.Sub(X, Y)
	fmt.Println("1-2 =", Z) //разность 1
	Z = Z.Sub(Y, X)
	fmt.Println("2-1 =", Z) // разность 2
	Z = Z.Mul(X, Y)
	fmt.Println("1*2 =", Z) // произведение
	Z = Z.Quo(X, Y)
	fmt.Println("1/2 =", Z) // частное 1
	Z = Z.Rem(X, Y)
	fmt.Println("1%2 =", Z) // остаток 1
	Z = Z.Quo(Y, X)
	fmt.Println("2/1 =", Z) // частное 2
	Z = Z.Rem(Y, X)
	fmt.Println("2%1 =", Z)        // остаток 2
	X1 := big.NewFloat(float64(x)) // переводим во float чтобы поделить с дробной частью
	Y1 := big.NewFloat(float64(y))
	Z1 := big.NewFloat(0.0).Quo(X1, Y1)
	fmt.Println("1/2 =", Z1) // отношение 1
	Z1 = Z1.Quo(Y1, X1)
	fmt.Println("2/1 =", Z1) // отношение 2
}
