package simple_and_midium

import (
	"fmt"
	"math/rand"
	"time"
)

type RandomizedSet struct {
	Nums []int
	flg map[int]int
	Len int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		Nums: make([]int, 128),
		Len: 0,
		flg: map[int]int{},
	}
}


func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.flg[val]; ok {
		return false
	}
	if this.Len >= len(this.Nums) {
		// o(n)
		this.Nums = append(this.Nums, val)
	} else {
		// o(1)
		this.Nums[this.Len] = val
	}
	this.flg[val] = this.Len
	this.Len++
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.flg[val]; !ok {
		return false
	}
	idx := this.flg[val]
	this.Len--
	this.Nums[idx] = this.Nums[this.Len]
	this.flg[this.Nums[this.Len]] = idx
	delete(this.flg, val)
	return true
}


func (this *RandomizedSet) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(this.Len)
	fmt.Println(this.Nums)
	fmt.Printf("random idx:%d, Len:%d\n", idx, this.Len)
	return this.Nums[idx]
}

func test()  {
	obj := Constructor()
	obj.Insert(1)
	obj.Remove(2)
	obj.Insert(2)
	r1 := obj.GetRandom()
	obj.Remove(1)
	obj.Insert(2)
	r2 := obj.GetRandom()
	fmt.Println(r1, r2)
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */