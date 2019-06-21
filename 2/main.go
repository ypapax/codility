package solution1

import (
	"fmt"
	"log"
)

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

var seasons = []string{"WINTER", "SPRING", "SUMMER", "AUTUMN"}

func Solution(A []int) string {
	if len(A)%len(seasons) != 0 {
		return fmt.Sprintf("error: tempretues array in to divisible by seasons amount: %d", len(seasons))
	}
	if len(A) == 0 {
		return fmt.Sprintf("error: empty temperatures array")
	}

	temperaturesInSeasonAmount := len(A) / len(seasons)
	log.Println("temperaturesInSeasonAmount", temperaturesInSeasonAmount)
	var highestAmplituted = 0
	var seasonWithHighestAltitudeNumber = 0
	var m = make([][]int, len(seasons))
	for i, t := range A {
		k := i / temperaturesInSeasonAmount
		log.Println("k", k)
		m[k] = append(m[k], t)
	}

	for seasonNumber, tt := range m {
		min, max := minAndMax(tt)
		ampl := max - min
		if ampl > highestAmplituted {
			highestAmplituted = ampl
			seasonWithHighestAltitudeNumber = seasonNumber
		}
	}
	return seasons[seasonWithHighestAltitudeNumber]
}

func minAndMax(A []int) (min, max int) {
	if len(A) == 0 {
		return
	}
	min = A[0]
	max = A[0]
	for _, t := range A {
		if t < min {
			min = t
		}
		if t > max {
			max = t
		}
	}
	return min, max
}
