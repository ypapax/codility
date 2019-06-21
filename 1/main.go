package solution1

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	// write your code in Go 1.4
	var countM = make(map[int]int)
	for _, v := range A {
		countM[v]++
	}
	var popularCount = 0
	var popularVal = 0
	for k, v := range countM {
		if v > popularCount {
			popularCount = v
			popularVal = k
		}
	}

	var totalRotations = 0
	for _, v := range A {
		if v == popularVal {
			continue
		}
		rotations := rotate(v, popularVal)
		totalRotations += rotations
	}
	return totalRotations
}

func rotate(from, to int) (rotations int) {
	if from+to == 7 {
		return 2
	}
	return 1
}
