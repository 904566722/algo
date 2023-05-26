package simple_and_midium

import "sort"

type ascInt []int

func (ai ascInt) Len() int {
	return len(ai)
}

func (ai ascInt) Less(i, j int) bool {
	return ai[i] < ai[j]
}

func (ai ascInt) Swap(i, j int) {
	ai[i], ai[j] = ai[j], ai[i]
}

type descInt []int

func (di descInt) Len() int {
	return len(di)
}

func (ai descInt) Less(i, j int) bool {
	return ai[i] > ai[j]
}

func (ai descInt) Swap(i, j int) {
	ai[i], ai[j] = ai[j], ai[i]
}

func isMonotonic(nums []int) bool {
	return sort.IsSorted(ascInt(nums)) || sort.IsSorted(descInt(nums))
}
