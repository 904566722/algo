package simple_and_midium

import (
	"fmt"
	"index/suffixarray"
	"reflect"
	"sort"
	"unsafe"
)

// s 的长度 l，对应的字串数量 l * (l+1) / 2
// 这种暴力方法会超时
func lastSubstring(s string) string {
	n := len(s)
	childS := make([]string, n*(n+1)/2)
	idx := 0
	for l := 1; l <= n; l++ {
		for i := 0; i+l <= n; i++ {
			//childS = append(childS, s[i:i+l])
			childS[idx] = s[i : i+l]
			idx++
		}
	}
	sort.Strings(childS)
	return childS[len(childS)-1]
}

func lastSubstring2(s string) string {
	sa := (*struct {
		_  []byte
		sa []int32
	})(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
	return s[sa[len(sa)-1]:]
}

type Index struct {
	_  []byte
	sa []int32
}

func lastSubstring3(s string) string {
	index := suffixarray.New([]byte(s))
	idx := (*Index)(unsafe.Pointer(index))
	fmt.Println(idx.sa)
	return ""
}

// ** struct practice
type person struct {
	name string
	age  int32
}

func structPractice() {
	people := &person{
		name: "索隆不喝酒",
		age:  20,
	}

	fmt.Println(people)  // &{索隆不喝酒 20}
	fmt.Println(*people) // {索隆不喝酒 20}
	fmt.Println(&people) // 0x1400000e038

	fmt.Printf("[value]\tname:%s\tage:%d\n", people.name, people.age)
	// [value]	name:索隆不喝酒	age:20
	fmt.Printf("[value]\tname:%s\tage:%d\n", (*people).name, (*people).age)
	// [value]	name:索隆不喝酒	age:20
	fmt.Printf("[addr]\tname:%d\tage:%d\n", &people.name, &people.age)
	// [addr]	name:1374389584016	age:1374389584032
}

type T struct {
	a int
}

func (t T) M1() {
	t.a++
}

func (t *T) M2() {
	t.a++
}

func structPractice2() {
	var t1 T
	var t2 = &T{}

	fmt.Println(reflect.TypeOf(t1)) // T
	fmt.Println(reflect.TypeOf(t2)) // *T

	t1.M1()
	t1.M2() // (&t1).M2()

	t2.M1() // (*t2).M1()
	t2.M2()

	T{}.M1()
	(&T{}).M2()

}
