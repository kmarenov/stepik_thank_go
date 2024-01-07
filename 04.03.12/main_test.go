package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"time"
)

// начало решения

// asLegacyDate преобразует время в легаси-дату
func asLegacyDate(t time.Time) string {
	partFirst := t.Unix()
	partSecond := strings.TrimRight(fmt.Sprintf("%09d", t.Nanosecond()), "0")
	if partSecond == "" {
		partSecond = "0"
	}
	return fmt.Sprintf("%d.%s", partFirst, partSecond)
}

// parseLegacyDate преобразует легаси-дату во время.
// Возвращает ошибку, если легаси-дата некорректная.
func parseLegacyDate(d string) (time.Time, error) {
	parts := strings.Split(d, ".")
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return time.Time{}, errors.New("wrong format")
	}
	partFirst, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	partSecond, err := strconv.ParseInt(addZeros(parts[1]), 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(partFirst, partSecond), nil
}

func addZeros(str string) string {
	if len(str) < 9 {
		diff := 9 - len(str)
		str = str + strings.Repeat("0", diff)
	}
	return str
}

// конец решения

func Test_asLegacyDate(t *testing.T) {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC): "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):         "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):         "0.0",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}

func Test_parseLegacyDate(t *testing.T) {
	samples := map[string]time.Time{
		"3600.123456789": time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":         time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":            time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":    time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
	}
	for src, want := range samples {
		got, err := parseLegacyDate(src)
		if err != nil {
			t.Fatalf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			t.Fatalf("%v: got %v, want %v", src, got, want)
		}
	}
}
