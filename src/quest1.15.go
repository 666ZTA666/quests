package main

import (
	"fmt"
	"math/rand"
	"strings"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как
//это исправить? Приведите корректный пример реализации.
//var justString string
//func someFunc() {
//v := createHugeString(1 << 10)
//justString = v[:100]
//}
//func main() {
//someFunc()
//}
//Во первых использование глобальных переменных в таком контексте — не правильно.
//В этой переменной может быть другое и, что важно, используемое другой функцией, значение,
// а мы при каждом вызове нашей функции будем ее переписывать.
//Во вторых если используются символы не из кодировки utf8, то создание подобного среза повлечет за собой
// взятие не 100 символов, а 100 байт, что в рунах может быть куда меньше 100 символов
//В целом не очень понятно зачем нам отрезать первые сто символов от строки, которую мы сами же и генерируем)
//Но если звёзды уже так сложились, то лучше взять сто рун, и не использовать глобальную переменную.
//Либо, если нам нужны 100 первых байт из строки, то сохранять их лучше в срез байт, хотя это не принципиально.
// И ещё кое-что, изменяя v(allStrings) с которой сделан срез, мы будем менять justString.

func createHugeString(size int) string {
	hugeString := strings.Builder{}
	for i := 0; i < size; i++ {
		hugeString.WriteRune('a' + rune(rand.Intn('я'-'А'+1)))
		// генератор строк, через strings.Builder
	}
	return hugeString.String()
}

func someFunc() (justString string) {
	//return перемеенную объявили прям в фукнции,
	// либо можно было указатель на строку передать,
	//если нам в конкретную переменную запись нужна
	allString := createHugeString(1 << 10)                            // это бывшая переменная V, так понятнее что она означает
	justString1 := allString[:100]                                    // тут первые 100 байт в строку
	justString = string(append([]rune{}, []rune(allString)[:100]...)) //а тут нам нужно первых 100 сиволов-рун
	// tmp := make([]rune, 100)
	// copy(tmp, []rune(allString)[:100])
	// justString = string(tmp)
	//Многострочный аналог.

	fmt.Println("justString =", justString)
	fmt.Println("len(justString)", len([]rune(justString)), "runs", len([]byte(justString)), "bytes")
	fmt.Println("cap(justString)", cap([]rune(justString)), "runs", cap([]byte(justString)), "bytes")
	// итогом получаем что justString независимая строка содержащая первые 100 рун
	fmt.Println("justString1 =", justString1)
	fmt.Println("len(justString1)", len([]rune(justString1)), "runs", len([]byte(justString1)), "bytes")
	fmt.Println("cap(justString1)", cap([]rune(justString1)), "runs", cap([]byte(justString1)), "bytes")
	// justString1 содержит первые 100 байт
	fmt.Println("allString =", allString)
	fmt.Println("len(allString)", len([]rune(allString)), "runs", len([]byte(allString)), "bytes")
	fmt.Println("cap(allString)", cap([]rune(allString)), "runs", cap([]byte(allString)), "bytes")
	// а это ради сравнения полная строка сгенерированная функцией createHugeString
	return
}
func main() {
	someFunc()
}
