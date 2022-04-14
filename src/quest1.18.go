package main

import (
	"fmt"
	"sync"
)

// Counter - counter with RWmutex for parallel counting
type Counter struct {
	sync.RWMutex
	counter int
}

// NewCounter - constructor for counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc - incrementing counter by 1
func (c *Counter) Inc() {
	c.Lock()
	c.counter++
	c.Unlock()
}

// Get - return value of counter
func (c *Counter) Get() int {
	c.RLock()
	defer c.RUnlock()
	s := c.counter
	return s
}

func main() {
	//конкретно в данном примере RWmutex не нужен, так как конкурентного чтения счётчика нет.
	c := NewCounter()
	wg := new(sync.WaitGroup)

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			c.Inc()
		}(wg)
	}

	wg.Wait()
	fmt.Println(c.Get())
}
