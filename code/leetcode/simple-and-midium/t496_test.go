package simple_and_midium

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_nextGreaterElement(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			args: args{
				nums1: []int{4, 1, 2},
				nums2: []int{1, 3, 4, 2},
			},
			want: []int{-1, 3, -1},
		},
		{
			args: args{
				nums1: []int{2, 4},
				nums2: []int{1, 2, 3, 4},
			},
			want: []int{3, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextGreaterElement(tt.args.nums1, tt.args.nums2), "nextGreaterElement(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}

func Test_myStack_pop(t *testing.T) {
	type fields struct {
		Nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			fields: fields{
				Nums: []int{1, 2, 3, 4},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &myStack{
				Nums: tt.fields.Nums,
			}
			assert.Equalf(t, tt.want, s.pop(), "pop()")
		})
	}
}

func Test_myStack_push(t *testing.T) {
	type fields struct {
		Nums []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &myStack{
				Nums: tt.fields.Nums,
			}
			s.push(tt.args.n)
		})
	}
}

func Test_myStack_top(t *testing.T) {
	type fields struct {
		Nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &myStack{
				Nums: tt.fields.Nums,
			}
			assert.Equalf(t, tt.want, s.top(), "top()")
		})
	}
}

func Test_nextGreaterElement1(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextGreaterElement(tt.args.nums1, tt.args.nums2), "nextGreaterElement(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}

func Test_myStack_isEmpty(t *testing.T) {
	type fields struct {
		Nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &myStack{
				Nums: tt.fields.Nums,
			}
			assert.Equalf(t, tt.want, s.isEmpty(), "isEmpty()")
		})
	}
}

func Test_myStack_pop1(t *testing.T) {
	type fields struct {
		Nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &myStack{
				Nums: tt.fields.Nums,
			}
			assert.Equalf(t, tt.want, s.pop(), "pop()")
		})
	}
}

func Test_myStack_push1(t *testing.T) {
	type fields struct {
		Nums []int
	}
	type args struct {
		n int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &myStack{
				Nums: tt.fields.Nums,
			}
			s.push(tt.args.n)
		})
	}
}

func Test_myStack_top1(t *testing.T) {
	type fields struct {
		Nums []int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &myStack{
				Nums: tt.fields.Nums,
			}
			assert.Equalf(t, tt.want, s.top(), "top()")
		})
	}
}

func Test_nextGreaterElement2(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextGreaterElement(tt.args.nums1, tt.args.nums2), "nextGreaterElement(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}

func Test_nextGreaterElement21(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//{
		//	args: args{
		//		nums1: []int{4, 1, 2},
		//		nums2: []int{1, 3, 4, 2},
		//	},
		//	want: []int{-1, 3, -1},
		//},
		//{
		//	args: args{
		//		nums1: []int{2, 4},
		//		nums2: []int{1, 2, 3, 4},
		//	},
		//	want: []int{3, -1},
		//},
		{
			args: args{
				nums1: []int{137, 59, 92, 122, 52, 131, 79, 236, 94},
				nums2: []int{137, 59, 92, 122, 52, 131, 79, 236, 94},
			},
			want: []int{236, 92, 122, 131, 131, 236, 236, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, nextGreaterElement2(tt.args.nums1, tt.args.nums2), "nextGreaterElement2(%v, %v)", tt.args.nums1, tt.args.nums2)
		})
	}
}
