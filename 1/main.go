package solution1

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

const diceSum = 7

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	// write your code in Go 1.4
	var unique = make(map[int]int)
	for _, v := range A {
		unique[v]++
	}
	var minRot = -1

	for popularVal := range unique {
		var totalRotations = 0
		for _, v := range A {
			if v == popularVal {
				continue
			}
			rotations := rotate(v, popularVal)
			totalRotations += rotations
		}
		if minRot < 0 {
			minRot = totalRotations
			continue
		}
		if totalRotations < minRot {
			minRot = totalRotations
		}
	}
	return minRot
}

func rotate(from, to int) (rotations int) {
	if from+to == diceSum {
		return 2
	}
	return 1
}
