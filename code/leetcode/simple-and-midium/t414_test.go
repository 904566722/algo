package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThirdMax(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, thirdMax([]int{1}))
	a.Equal(2, thirdMax([]int{1, 1, 2}))
	a.Equal(1, thirdMax([]int{1, 0}))
	a.Equal(6, thirdMax([]int{10, 9, 9, 9, 9, 6, 3, 4, 2, 2, 2}))

}
