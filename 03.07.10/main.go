package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	lock sync.RWMutex
	vals map[string]int
}

func (c *Counter) Increment(str string) {
	c.lock.Lock()
	c.vals[str]++
	c.lock.Unlock()
}

func (c *Counter) Value(str string) int {
	c.lock.RLock()
	v := c.vals[str]
	c.lock.RUnlock()
	return v
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.lock.RLock()
	for k, v := range c.vals {
		fn(k, v)
	}
	c.lock.RUnlock()
}

func NewCounter() *Counter {
	return &Counter{
		lock: sync.RWMutex{},
		vals: make(map[string]int),
	}
}

// конец решения

func main() {
	counter := NewCounter()

	var wg sync.WaitGroup
	wg.Add(3)

	increment := func(key string, val int) {
		defer wg.Done()
		for ; val > 0; val-- {
			counter.Increment(key)
		}
	}

	go increment("one", 100)
	go increment("two", 200)
	go increment("three", 300)

	wg.Wait()

	fmt.Println("two:", counter.Value("two"))

	fmt.Print("{ ")
	counter.Range(func(key string, val int) {
		fmt.Printf("%s:%d ", key, val)
	})
	fmt.Println("}")
}
