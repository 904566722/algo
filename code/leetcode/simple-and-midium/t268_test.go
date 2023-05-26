package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, missingNumber([]int{3, 0, 1}))
	a.Equal(1, missingNumber([]int{0}))
}

func TestMissingNumber2(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, missingNumber2([]int{3, 0, 1}))
	a.Equal(1, missingNumber2([]int{0}))
	a.Equal(0, missingNumber2([]int{1, 2}))
}

func TestMissingNumber3(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, missingNumber3([]int{3, 0, 1}))
	a.Equal(1, missingNumber3([]int{0}))
	a.Equal(0, missingNumber3([]int{1, 2}))
}
