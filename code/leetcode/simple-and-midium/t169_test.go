package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	a := assert.New(t)
	element1 := majorityElement([]int{3, 2, 3})
	element2 := majorityElement([]int{3})
	element3 := majorityElement([]int{3, 3, 1, 3})
	a.Equal(3, element1)
	a.Equal(3, element2)
	a.Equal(3, element3)
}

func TestMajorityElement2(t *testing.T) {
	a := assert.New(t)
	element1 := majorityElement2([]int{3, 2, 3})
	element2 := majorityElement2([]int{3})
	element3 := majorityElement2([]int{3, 3, 1, 3})
	a.Equal(3, element1)
	a.Equal(3, element2)
	a.Equal(3, element3)
}
