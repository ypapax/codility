package solution1

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	as := assert.New(t)
	inp := []int{6, 1, 4, 6, 3, 2, 7, 4}
	k, l := 3, 2
	exp := 24
	act := Solution(inp, k, l)
	as.Equal(exp, act)
}

func TestSolution2(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	as := assert.New(t)
	inp := []int{10, 19, 15}
	k, l := 2, 2
	exp := -1
	act := Solution(inp, k, l)
	as.Equal(exp, act)
}
