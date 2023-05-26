package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPoisonedDuration(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, findPoisonedDuration([]int{1, 4}, 2))
	a.Equal(3, findPoisonedDuration([]int{1, 2}, 2))
	a.Equal(4, findPoisonedDuration([]int{1}, 4))
}
