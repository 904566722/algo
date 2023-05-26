package main

import (
	"fmt"
	"leetcode/simple-and-midium"
)

func main() {
	fmt.Println(simple_and_midium.FindCircleNum([][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1},
	}))
}

func testFor() {
	cnt := 0
	for i := 0; i < 10; i++ {
		cnt++
	}
	return
}
