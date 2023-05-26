package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	a.Equal(true, containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	a.Equal(false, containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}

func TestContainsNearbyDuplicate2(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, containsNearbyDuplicate2([]int{1, 2, 3, 1}, 3))
	a.Equal(true, containsNearbyDuplicate2([]int{1, 0, 1, 1}, 1))
	a.Equal(false, containsNearbyDuplicate2([]int{1, 2, 3, 1, 2, 3}, 2))
}
