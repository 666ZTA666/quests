package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func main() {
	var u = int64(math.MaxInt64) * -1
	fmt.Printf("Начальное значение числа =%v\nИли %b\n\n", u, u)
	//n := "0"
	//sn := n[0] //"0" = 48, "1" = 49, "-" = 45
	for {
		BitFunc(&u)
	}
}

// работает как часы, но из-за большого количества обработок строк итп тратит иногда чуть больше времени
func mytjFunc(u *int64) *int64 {
	var (
		c int64
		s string
		z bool
		i int
	)
	if u == nil {
		fmt.Println("u == nil")
		return nil
	}
	fmt.Println("Какой по очередности бит поменять? от 1 до 64, где 1 отвечает за значения 0\\1 в 10, а 64 за знак")
	_, err := fmt.Scanln(&i)
	start := time.Now()
	i = 64 - i
	if err != nil {
		fmt.Println(err)
		return u
	}
	if i >= 64 || i < 0 {
		fmt.Println("принимаются числа только от 1 до 64")
		return u
	}
	if i == 0 && *u != 0 {

	} else if i == 0 && *u == 0 {
		u = u
	} else {
		c = 1
		z = *u >= 0
		s = fmt.Sprintf("%b", u)
		if len(s) < 64 && z {
			for len(s) < 64 {
				s = fmt.Sprint(0, s)
			}
		} else if len(s) < 64 && !z {
			s = strings.Trim(s, "-")
			for len(s) < 63 {
				s = fmt.Sprint(0, s)
			}
			s = fmt.Sprint("-", s)
		}
		for j := 1; j < (64 - i); j++ {
			c *= 2
		}
		if s[i] == 48 && z == true || s[i] == 49 && z == false {
			*u += c
		} else {
			*u -= c
		}
	}
	fmt.Printf("%b\n%v\n", u, u)
	fmt.Println("время работы функции =", time.Now().Sub(start))
	return u
}

// работает не совсем идеально, но должна управляться побыстрее
func BitFunc(u *int64) *int64 {
	var (
		i int
	)
	fmt.Println("Какой по очередности бит поменять? от 1 до 64, где 1 отвечает за значения 0\\1 в 10, а 64 за знак")
	_, err := fmt.Scanln(&i)
	start := time.Now()
	if err != nil {
		fmt.Println(err)
		return u
	}
	if i > 64 || i <= 0 {
		fmt.Println("принимаются числа только от 1 до 64")
		return u
	}
	if i == 64 {
		*u = -*u
		fmt.Printf("%b\n%v\n", *u, *u)
		fmt.Println(time.Now().Sub(start))
		return u
	}
	*u ^= 1 << (i - 1)
	fmt.Printf("%b\n%v\n", *u, *u)
	fmt.Println(time.Now().Sub(start))
	return u
}

// Была идея полностью на срезах байт сделать, чтобы каждое число хранилось как срез байт, но стало оч сложно может todo
