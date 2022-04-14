package main

import "fmt"

type client struct{}

// ClientFunc
// Есть некая клиент-функция, которая принимает только объекты первого интерфейса.
func (c *client) ClientFunc(com First) {
	fmt.Println("Client use First interfaced Staff")
	com.FirstFunc()
}

// First
// Вот собственно первый интерфейс. Он содержит свои методы.
type First interface {
	FirstFunc()
}
type FirstStruct struct{}

func (m *FirstStruct) FirstFunc() {
	fmt.Println("FirstFunc")
}

// Second
// Вот интерфейс номер два, у него свои методы, которые не подходят интерфейсу №1
// Но нам необходимо, чтобы клиент мог использовать наши методы.
type Second interface {
	SecondFunc()
}
type SecondStruct struct{}

func (w *SecondStruct) SecondFunc() {
	fmt.Println("SecondFunc")
}

// Adapter
// Вот адаптер, у него есть функция первого интерфейса, а за счет композиции имеет возможность
// вызывать методы второго интерфейса. Что позволяет нам передать его в качестве аргумента
// клиент-функции и использовать там методы второго интерфейса. При этом не были переписаны методы
// первого и второго интерфейсов.
type Adapter struct {
	SecondStruct *SecondStruct
}

func (w *Adapter) FirstFunc() {
	fmt.Println("Adapter converts First to Second")
	w.SecondStruct.SecondFunc()
}

func main() {
	client := &client{}
	First := &FirstStruct{}
	fmt.Println("Client use First")
	client.ClientFunc(First)

	SecondSt := &SecondStruct{}
	Adapter := &Adapter{SecondStruct: SecondSt}
	fmt.Println("Client use Adapted second")
	client.ClientFunc(Adapter)
}

/* Итого: не переписывая метод первой и второй структуры, а также не меняя содержимое первого и второго интерфейсов,
Мы смогли передать в клиентскую функцию адаптированную вторую структуру с её методами.*/
