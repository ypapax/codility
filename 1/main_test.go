package solution1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	as := assert.New(t)
	inp := []int{1, 2, 3}
	exp := 2

	act := Solution(inp)
	as.Equal(exp, act)
}

func TestSolution2(t *testing.T) {
	as := assert.New(t)
	inp := []int{1, 1, 6}
	exp := 2

	act := Solution(inp)
	as.Equal(exp, act)
}

func TestSolution3(t *testing.T) {
	as := assert.New(t)
	inp := []int{1, 6, 2, 3}
	exp := 3

	act := Solution(inp)
	as.Equal(exp, act)
}
