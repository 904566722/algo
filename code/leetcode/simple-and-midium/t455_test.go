package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test455(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	a.Equal(2, findContentChildren([]int{1, 2, 3}, []int{1, 2}))
	a.Equal(2, findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	a.Equal(1, findContentChildren([]int{1, 2}, []int{3}))
}
