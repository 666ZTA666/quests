package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MyComparable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool
}

// мой контракт для использования в функции, решил попользовать дженерики, вместо пустого интерфейса)

func main() {
	//Stackoverflow говорит следующее:
	//Simple Intersection: Compare each element in A to each in B (O(n^2)).
	//Hash Intersection: Put them into a hash table (O(n*x)),
	// where x is a factor of hash function efficiency (between 1 and 2).
	//Sorted Intersection: Sort A and do an optimized intersection (O(n*log(n))).
	// чуток посчитав и подумав выбираю вариант с хэшем).
	//Вариант с сортировкой работает лучше в когда (n) размер последовательностей до 7 элементов.

	var n1, n2, max int
	for {
		fmt.Println("введите количество элементов в первом слайсе")
		_, err := fmt.Scanln(&n1) // считали, записали 1
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n1 > 0 {
			break
		}
	}
	for {
		fmt.Println("введите количество элементов во втором слайсе")
		_, err := fmt.Scanln(&n2) // считали, записали 2
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n2 > 0 {
			break
		}
	}
	for {
		fmt.Println("введите максимальное значение элемента в слайсе")
		_, err := fmt.Scanln(&max) // считали записали max
		if err != nil {
			fmt.Println(err)
			continue
		}
		if max > 0 {
			break
		}
	}

	if n1 > n2 {
		// do nothing
	} else {
		n1, n2 = n2, n1
	}
	// тут было чуток удобнее отталкиваться от размеров срезов, и хэшить меньший, а сравнивать с большим
	var SliseInt1 = make([]int, n1)
	var SliseInt2 = make([]int, n2)
	for i := 0; i < n1; i++ {
		SliseInt1[i] = rand.Intn(max) // заполняем слайсы псевдослучайными числами
	}
	for i := 0; i < n2; i++ {
		SliseInt2[i] = rand.Intn(max)
		//В go нет генератора случайных строк (но чуть попозже Я его напишу, было лень переносить сюда),
		//поэтому заполняю случайными int'ами.
	}
	Res, u, t := Hash(SliseInt2, SliseInt1)
	// собственно функция которая возвращает срез, который отражает пересечение двух множеств
	// количество элементов в этом срезе и время работы.
	fmt.Println(Res)
	fmt.Println("Размер большего массива:", n1, "размер меньшего массива:", n2, "\nРазброс значений элементов массива:", max, "найдено совпадений:", u, "время работы функции:", t)
	// выводим все данные
}

func Hash[T MyComparable](a []T, b []T) ([]T, int, time.Duration) {
	// засекаем время
	start := time.Now()
	// создадим срез результатов
	set := make([]T, 0)
	// и карту для хэширования данных
	hash := make(map[T]struct{})
	for i := 0; i < len(a); i++ {
		hash[a[i]] = struct{}{} // записываем в мапу все элементы меньшего среза
	}
	for i := 0; i < len(b); i++ { // проходим в цикле по всем элементам большего среза
		if _, found := hash[b[i]]; found {
			set = append(set, b[i])
		}
	}
	// возвращаем срез, длину и время
	// по факту время можно было просто вывести и не возвращать
	// длина массива нужна, чтобы проверить её на 0, и в случае нуля не работать с пустым срезом, ну это на будущее...
	return set, len(set), time.Now().Sub(start)
}
