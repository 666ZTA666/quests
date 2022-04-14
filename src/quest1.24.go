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
		s = readLineNow(w)        // читаем строку через bufio
		arrs := strings.Fields(s) // делим на слова
		if len(arrs) < 2 {
			fmt.Println("мало значений") // проверка
			continue
		}
		x1, err = strconv.ParseInt(arrs[0], 10, 64) // парсим слово на int
		if err != nil {
			fmt.Println("что-то не так", err) // проверка
			continue
		}
		y1, err = strconv.ParseInt(arrs[1], 10, 64) // парсим слово на int
		if err != nil {
			fmt.Println("что-то не так", err) // проверка
			continue
		}
		break // если всё ок выходим
	}
	for {
		fmt.Println("введите координаты второй точки, желательно через пробел")
		s = readLineNow(w)        // читаем строку через bufio
		arrs := strings.Fields(s) // делим на слова
		if len(arrs) < 2 {
			fmt.Println("мало значений") // проверка
			continue
		}
		x2, err = strconv.ParseInt(arrs[0], 10, 64) // парсим слово на int
		if err != nil {
			fmt.Println("что-то не так", err) // проверка
			continue
		}
		y2, err = strconv.ParseInt(arrs[1], 10, 64) // парсим слово на int
		if err != nil {
			fmt.Println("что-то не так", err) // проверка
			continue
		}
		break
	}
	// создаем точки по координатам
	a := NewPoint(x1, y1)
	b := NewPoint(x2, y2)
	fmt.Println(a.Dist(b)) // выводим дистанцию
}

// Point сама структура
type Point struct {
	x, y int64
}

// конструктор для точек
func NewPoint(x int64, y int64) *Point {
	return &Point{x: x, y: y}
}

//поиск дистанции для одной точки относительно другой
func (p *Point) Dist(po *Point) float64 {
	distX := dist(p.x, po.x)
	distY := dist(p.y, po.y)
	return math.Sqrt(float64(distX*distX + distY*distY))
}

//возвращает положительное значение растояния между точками на одной оси
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

// читаем строку удаляя все лишние  непотребства сразу
func readLineNow(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
