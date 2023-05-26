package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	a := assert.New(t)
	a.Equal([]string{"0->2", "4->5", "7"}, summaryRanges([]int{0, 1, 2, 4, 5, 7}))
	a.Equal([]string{"0", "2->4", "6", "8->9"}, summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
	a.Equal([]string{"0"}, summaryRanges([]int{0}))
	a.Equal([]string{"0->1"}, summaryRanges([]int{0, 1}))
	a.Equal([]string{}, summaryRanges([]int{}))
}
