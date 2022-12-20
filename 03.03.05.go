package main

import (
	"fmt"
	"time"
)

func await(fn func() any) any {
    done := make(chan any, 1)    // (1)
    go func() {
        done <- fn()             // (2)
    }()
    return <-done
}

// gather выполняет переданные функции одновременно
// и возвращает срез с результатами, когда они готовы
func gather(funcs []func() any) []any {
	// начало решения

	// выполните все переданные функции,
	// соберите результаты в срез
	// и верните его

	type pair struct {
		idx int
		val any
	}
	
	done := make(chan pair)
	for i, f := range funcs {
		go func(i int, fn func() any) {
			done <- pair{i, fn()}
		}(i, f)
	}

	results := make([]any, len(funcs))
	for i := 0; i < len(funcs); i++ {
		r := <-done
		results[r.idx] = r.val
	}
	close(done)

	return results

	// конец решения
}

// squared возвращает функцию,
// которая считает квадрат n
func squared(n int) func() any {
	return func() any {
		time.Sleep(time.Duration(n) * 100 * time.Millisecond)
		return n * n
	}
}

func main() {
	funcs := []func() any{squared(2), squared(3), squared(4)}

	start := time.Now()
	nums := gather(funcs)
	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Println(nums)
	fmt.Printf("Took %.0f ms\n", elapsed)
}
