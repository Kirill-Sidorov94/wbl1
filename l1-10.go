package main

import (
	"fmt"
	"math"
)

func temperatureGrouping() {
	sl := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	m := make(map[float64][]float64, len(sl))

	for i := range sl {
		var key float64
		if sl[i] > 0 {
			key = roundPositiveNum(sl[i])
		} else {
			key = roundNegativeNum(sl[i])
		}

		m[key] = append(m[key], sl[i])
	}

	fmt.Println(m)
}

func roundPositiveNum(num float64) float64 {
	return math.Floor(math.Floor(math.Trunc(num)) / 10) * 10
}

func roundNegativeNum(num float64) float64 {
	return math.Floor(math.Ceil(math.Trunc(num)) / 10) * 10
}
