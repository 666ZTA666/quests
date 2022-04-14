package main

import (
	"fmt"
	"math/big"
)

func main() {
	var x int64
	for {
		fmt.Println("введите первое число")
		_, err := fmt.Scanln(&x)
		if err != nil {
			fmt.Println(err)
		}
		if x != 0 {
			break
		} else {
			fmt.Println("Делить на ноль мы не будем")
		}
	}
	var y int64
	for {
		fmt.Println("введите второе число")
		_, err := fmt.Scanln(&y)
		if err != nil {
			fmt.Println(err)
		}
		if y != 0 {
			break
		} else {
			fmt.Println("Делить на ноль мы не будем")
		}
	}

	X := big.NewInt(x)
	Y := big.NewInt(y)
	fmt.Println("1 =", X, "2 =", Y)

	Z := big.NewInt(0).Add(Y, X)
	fmt.Println("1+2 =", Z)
	Z = Z.Sub(X, Y)
	fmt.Println("1-2 =", Z)
	Z = Z.Sub(Y, X)
	fmt.Println("2-1 =", Z)
	Z = Z.Mul(X, Y)
	fmt.Println("1*2 =", Z)
	Z = Z.Quo(X, Y)
	fmt.Println("1/2 =", Z)
	Z = Z.Rem(X, Y)
	fmt.Println("1%2 =", Z)
	Z = Z.Quo(Y, X)
	fmt.Println("2/1 =", Z)
	Z = Z.Rem(Y, X)
	fmt.Println("2%1 =", Z)
	X1 := big.NewFloat(float64(x))
	Y1 := big.NewFloat(float64(y))
	Z1 := big.NewFloat(0.0).Quo(X1, Y1)
	fmt.Println("1/2 =", Z1)
	Z1 = Z1.Quo(Y1, X1)
	fmt.Println("2/1 =", Z1)
}
