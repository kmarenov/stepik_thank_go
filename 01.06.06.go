package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
)

// element - интерфейс элемента последовательности
type element interface{}

// weightFunc - функция, которая возвращает вес элемента
type weightFunc func(element) int

// iterator - интерфейс, который умеет
// поэлементно перебирать последовательность
type iterator interface {
    next() bool
	val() int
}

// intIterator - итератор по целым числам
// (реализует интерфейс iterator)
type intIterator struct {
	vals []int
	currIdx int
	first bool
}

func (i *intIterator) next() bool {
	if !i.first {
		i.currIdx++
	}
	i.first = false
	return i.currIdx < len(i.vals)
}

func (i *intIterator) val() int {
	return i.vals[i.currIdx]
}


// методы intIterator, которые реализуют интерфейс iterator

// конструктор intIterator
func newIntIterator(src []int) *intIterator {
    return &intIterator{src, 0, true}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// main находит максимальное число из переданных на вход программы.
func main() {
    nums := readInput()
    it := newIntIterator(nums)
    weight := func(el element) int {
        return el.(int)
    }
    m := max(it, weight)
    fmt.Println(m)
}

// max возвращает максимальный элемент в последовательности.
// Для сравнения элементов используется вес, который возвращает
// функция weight.
func max(it iterator, weight weightFunc) element {
    var maxEl element
    for it.next() {
        curr := it.val()
        if maxEl == nil || weight(curr) > weight(maxEl) {
            maxEl = curr
        }
    }
    return maxEl
}

// readInput считывает последовательность целых чисел из os.Stdin.
func readInput() []int {
    var nums []int
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        num, err := strconv.Atoi(scanner.Text())
        if err != nil {
            log.Fatal(err)
        }
        nums = append(nums, num)
    }
    return nums
}