package solution1

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	as := assert.New(t)
	inp := []int{-3, -14, -5, 7, 8, 42, 8, 3}
	exp := "SUMMER"

	act := Solution(inp)
	as.Equal(exp, act)
}

func TestSolution2(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	as := assert.New(t)
	inp := []int{2, -3, 3, 1, 10, 8, 2, 5, 13, -5, 3, -18}
	exp := "AUTUMN"

	act := Solution(inp)
	as.Equal(exp, act)
}
