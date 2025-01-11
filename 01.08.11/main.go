package main

import "fmt"

// начало решения

// Map - карта "ключ-значение".
type Map[K comparable, V any] struct {
	m map[K]V
}

// Set устанавливает значение для ключа.
func (s *Map[K, V]) Set(key K, val V) {
	if s.m == nil {
		s.m = make(map[K]V)
	}

	s.m[key] = val
}

// Get возвращает значение по ключу.
func (s *Map[K, V]) Get(key K) V {
	if s.m == nil {
		s.m = make(map[K]V)
	}

	return s.m[key]
}

// Keys возвращает срез ключей карты.
// Порядок ключей неважен, и не обязан совпадать
// с порядком значений из метода Values.
func (s *Map[K, V]) Keys() []K {
	if s.m == nil {
		s.m = make(map[K]V)
	}

	keys := make([]K, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}

	return keys
}

// Values возвращает срез значений карты.
// Порядок значений неважен, и не обязан совпадать
// с порядком ключей из метода Keys.
func (s *Map[K, V]) Values() []V {
	if s.m == nil {
		s.m = make(map[K]V)
	}

	vals := make([]V, 0, len(s.m))
	for _, v := range s.m {
		vals = append(vals, v)
	}

	return vals
}

// конец решения

func main() {
	m := Map[string, int]{}
	m.Set("one", 1)
	m.Set("two", 2)

	fmt.Println(m.Get("one")) // 1
	fmt.Println(m.Get("two")) // 2

	fmt.Println(m.Keys())   // [one two]
	fmt.Println(m.Values()) // [1 2]
}
