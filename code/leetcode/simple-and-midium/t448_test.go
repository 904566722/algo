package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{5, 6}, findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	a.Equal([]int{2}, findDisappearedNumbers([]int{1, 1}))
}

func TestFindDisappearedNumbers2(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{5, 6}, findDisappearedNumbers2([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	a.Equal([]int{2}, findDisappearedNumbers2([]int{1, 1}))
}
