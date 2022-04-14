package main

import (
	"fmt"
	"sync"
)

// Counter - счетчик с рв мьютексом
type Counter struct {
	sync.RWMutex
	counter int
}

// NewCounter - конструктор счетчика
func NewCounter() *Counter {
	return &Counter{}
}

// Inc - увеличитель счетчика
func (c *Counter) Inc() {
	c.Lock()
	c.counter++
	c.Unlock()
}

// Get - возвращает значение счетчика
func (c *Counter) Get() int {
	c.RLock()
	s := c.counter
	c.RUnlock()
	return s
}

func main() {
	//конкретно в данном примере RWmutex не нужен, так как конкурентного чтения счётчика нет.
	c := NewCounter()
	// создали новый счетчик
	wg := new(sync.WaitGroup)
	// создали группу ожидания

	wg.Add(10000) // добавили 10000 членов группы ожидания
	for i := 0; i < 10000; i++ {
		// после цикла на 10 000 итераций счетчик должен быть равен 10 000
		go func(wg *sync.WaitGroup) {
			c.Inc()   // увеличили счетчик
			wg.Done() // убрали единичку из группы
		}(wg)
	}

	wg.Wait() // ждем завершения всех горутин
	// и выводим счетчик
	fmt.Println(c.Get())
}
