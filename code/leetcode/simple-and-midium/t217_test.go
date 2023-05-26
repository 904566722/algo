package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicate2(t *testing.T) {
	a := assert.New(t)
	a.Equal(false, containsDuplicate2([]int{1, 2, 3, 4, 5, 6}))
	a.Equal(true, containsDuplicate2([]int{1, 2, 3, 4, 5, 5}))
	a.Equal(false, containsDuplicate2([]int{1}))
	a.Equal(true, containsDuplicate2([]int{1, 1}))
}
