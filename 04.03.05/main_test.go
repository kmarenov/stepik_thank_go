package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

// начало решения

// TimeOfDay описывает время в пределах одного дня
type TimeOfDay struct {
	t time.Time
}

// Hour возвращает часы в пределах дня
func (t TimeOfDay) Hour() int {
	return t.t.Hour()
}

// Minute возвращает минуты в пределах часа
func (t TimeOfDay) Minute() int {
	return t.t.Minute()
}

// Second возвращает секунды в пределах минуты
func (t TimeOfDay) Second() int {
	return t.t.Second()
}

// String возвращает строковое представление времени
// в формате чч:мм:сс TZ (например, 12:34:56 UTC)
func (t TimeOfDay) String() string {
	return fmt.Sprintf("%02d:%02d:%02d %s", t.t.Hour(), t.t.Minute(), t.t.Second(), t.t.Location())
}

// Equal сравнивает одно время с другим.
// Если у t и other разные локации - возвращает false.
func (t TimeOfDay) Equal(other TimeOfDay) bool {
	if t.t.Location().String() != other.t.Location().String() {
		return false
	}

	return t.t.Equal(other.t)
}

// Before возвращает true, если время t предшествует other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.t.Location().String() != other.t.Location().String() {
		return false, errors.New("different location")
	}

	return t.t.Before(other.t), nil
}

// After возвращает true, если время t идет после other.
// Если у t и other разные локации - возвращает ошибку.
func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.t.Location().String() != other.t.Location().String() {
		return false, errors.New("different location")
	}

	return t.t.After(other.t), nil
}

// MakeTimeOfDay создает время в пределах дня
func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	return TimeOfDay{
		t: time.Date(2024, 1, 1, hour, min, sec, 0, loc),
	}
}

// конец решения

func Test(t *testing.T) {
	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	if t1.Equal(t2) {
		t.Errorf("%v should not be equal to %v", t1, t2)
	}

	before, _ := t1.Before(t2)
	if !before {
		t.Errorf("%v should be before %v", t1, t2)
	}

	after, _ := t1.After(t2)
	if after {
		t.Errorf("%v should NOT be after %v", t1, t2)
	}
}
