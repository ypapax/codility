package solution1

import "log"

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K, L int) int {
	// when Alice is to the left of Bob
	max1 := getMaxApples(A, K, L)
	// when Bob is to the left of Alice
	max2 := getMaxApples(A, L, K)
	if max1 > max2 {
		return max1
	}
	return max2

}

func getMaxApples(A []int, K, L int) int {
	var maxApples = -1
	if K <= 0 || L <= 0 {
		return maxApples
	}
	for firstTree := 0; firstTree < len(A); firstTree++ {
		var applesForAlice, applesForBob int
		var lastAliceTree = firstTree + K - 1
		if lastAliceTree >= len(A) {
			continue
		}
		for i := firstTree; i <= lastAliceTree; i++ {
			applesForAlice += A[i]
		}

		firstBobTree := lastAliceTree + 2
		lastBobTree := firstBobTree + L - 1
		if lastBobTree >= len(A) {
			continue
		}
		for i := firstBobTree; i <= lastBobTree; i++ {
			applesForBob += A[i]
		}

		totalApples := applesForAlice + applesForBob

		if totalApples > maxApples {
			maxApples = totalApples
		}
		log.Printf("totalApples %d for firstTree %d, lastAliceTree %d, firstBobTree %d, lastBobTree %d\n", totalApples, firstTree, lastAliceTree, firstBobTree, lastBobTree)
	}
	return maxApples
}
