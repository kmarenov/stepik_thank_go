package main

import (
	"fmt"
	"math/rand"
	"time"
)

// начало решения

func delay(dur time.Duration, fn func()) func() {
	done := make(chan struct{})
	start := time.Now()
	go func(dur time.Duration){
		defer func() {
			close(done) 
			done = nil
		}()
		for {
			select {
				case <-done: return
				default: {
					now := time.Now()
					if start.Add(dur).Before(now) {
						fn()
						return
					}
				}
			}
		}
	}(dur)
	cancel := func() {
		if done != nil {
			done<-struct{}{}
		}
	}
	return cancel
}

// конец решения

func main() {
	rand.Seed(time.Now().Unix())

	work := func() {
		fmt.Println("work done")
	}

	cancel := delay(100*time.Millisecond, work)

	time.Sleep(10 * time.Millisecond)
	if rand.Float32() < 0.5 {
		cancel()
		fmt.Println("delayed function canceled")
	}
	time.Sleep(100 * time.Millisecond)
}
