package main

import (
	"strconv"
	"strings"
	"testing"
)

// начало решения

// calcDistance возвращает общую длину маршрута в метрах
func calcDistance(directions []string) int {
	d := 0
	for _, direction := range directions {
		for _, command := range strings.Fields(direction) {
			d += parseKilometers(command)
			d += parseMeters(command)
		}
	}
	return d
}

func parseKilometers(src string) int {
	dist := 0
	if strings.HasSuffix(src, "km") {
		d := strings.TrimRight(src, "km")
		f, err := strconv.ParseFloat(d, 64)
		if err == nil {
			dist = int(f * 1000.0)
		}
	}
	return dist
}

func parseMeters(src string) int {
	dist := 0
	if strings.HasSuffix(src, "m") {
		d := strings.TrimRight(src, "m")
		f, err := strconv.Atoi(d)
		if err == nil {
			dist = f
		}
	}
	return dist
}

// конец решения

func Test(t *testing.T) {
	directions := []string{
		"100m to intersection",
		"turn right",
		"straight 300m",
		"enter motorway",
		"straight 5km",
		"exit motorway",
		"500m straight",
		"turn sharp left",
		"continue 100m to destination",
	}
	const want = 6000
	got := calcDistance(directions)
	if got != want {
		t.Errorf("%v: got %v, want %v", directions, got, want)
	}
}
