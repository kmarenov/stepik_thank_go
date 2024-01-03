package main

import (
	"fmt"
	"sync"
)

// начало решения

type Counter struct {
	storage map[string]int
	mu      *sync.Mutex
}

func (c *Counter) Increment(str string) {
	c.mu.Lock()
	c.storage[str]++
	c.mu.Unlock()
}

func (c *Counter) Value(str string) int {
	c.mu.Lock()
	val := c.storage[str]
	c.mu.Unlock()
	return val
}

func (c *Counter) Range(fn func(key string, val int)) {
	c.mu.Lock()
	for k, v := range c.storage {
		fn(k, v)
	}
	c.mu.Unlock()
}

func NewCounter() *Counter {
	return &Counter{
		storage: make(map[string]int),
		mu:      &sync.Mutex{},
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
