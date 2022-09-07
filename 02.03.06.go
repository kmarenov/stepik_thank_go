package main

// не удаляйте импорты, они используются при проверке
import (
    "fmt"
    "math/rand"
    "os"
    "testing"
)

// реализуйте быстрое множество
type IntSet struct {
    elems *map[int]struct{}
}

// MakeIntSet creates an empty set.
func MakeIntSet() IntSet {
    elems := make(map[int]struct{})
    return IntSet{&elems}
}

// Contains reports whether an element is within the set.
func (s IntSet) Contains(elem int) bool {
    _, ok := (*s.elems)[elem]
    return ok
}

// Add adds an element to the set.
func (s IntSet) Add(elem int) bool {
    if s.Contains(elem) {
        return false
    }
    (*s.elems)[elem] = struct{}{}
    return true
}
