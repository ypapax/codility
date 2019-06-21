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
