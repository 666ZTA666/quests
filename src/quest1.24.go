package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Разработать программу нахождения расстояния между двумя точками, которые
	//представлены в виде структуры Point с инкапсулированными параметрами x,y и
	//конструктором.

	w := bufio.NewReader(os.Stdin)
	var (
		err            error
		x1, x2, y1, y2 int64
		s              string
	)
	for {
		fmt.Println("введите координаты первой точки, желательно через пробел")
		s = readLineNow(w)
		arrs := strings.Fields(s)
		if len(arrs) < 2 {
			fmt.Println("мало значений")
			continue
		}
		x1, err = strconv.ParseInt(arrs[0], 10, 64)
		if err != nil {
			fmt.Println("что-то не так", err)
			continue
		}
		y1, err = strconv.ParseInt(arrs[1], 10, 64)
		if err != nil {
			fmt.Println("что-то не так", err)
			continue
		}
		break
	}
	for {
		fmt.Println("введите координаты второй точки, желательно через пробел")
		s = readLineNow(w)
		arrs := strings.Fields(s)
		if len(arrs) < 2 {
			fmt.Println("мало значений")
			continue
		}
		x2, err = strconv.ParseInt(arrs[0], 10, 64)
		if err != nil {
			fmt.Println("что-то не так", err)
			continue
		}
		y2, err = strconv.ParseInt(arrs[1], 10, 64)
		if err != nil {
			fmt.Println("что-то не так", err)
			continue
		}
		break
	}
	a := NewPoint(x1, y1)
	b := NewPoint(x2, y2)
	fmt.Println(a.Dist(b))
}

type Point struct {
	x, y int64
}

func NewPoint(x int64, y int64) *Point {
	return &Point{x: x, y: y}
}
func (p *Point) Dist(po *Point) float64 {
	distX := dist(p.x, po.x)
	distY := dist(p.y, po.y)
	return math.Sqrt(float64(distX*distX + distY*distY))
}
func dist(x1, x2 int64) int64 {
	switch x1 > x2 {
	case true:
		return x1 - x2
	case false:
		return x2 - x1
	default:
		return 0
	}
}
func readLineNow(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
