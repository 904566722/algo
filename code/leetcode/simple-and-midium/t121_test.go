package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	assert := assert.New(t)
	profit1 := maxProfit([]int{7, 1, 5, 3, 6, 4})
	profit2 := maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(5, profit1)
	assert.Equal(0, profit2)
}

func TestMaxProfit2(t *testing.T) {
	assert := assert.New(t)
	profit1 := maxProfit2([]int{7, 1, 5, 3, 6, 4})
	profit2 := maxProfit2([]int{7, 6, 4, 3, 1})
	profit3 := maxProfit2([]int{2, 5, 1, 3})
	assert.Equal(5, profit1)
	assert.Equal(0, profit2)
	assert.Equal(3, profit3)
}
