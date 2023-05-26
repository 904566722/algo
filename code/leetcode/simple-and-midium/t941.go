package simple_and_midium

// 用一个标志来表示是上坡还是下坡，flg = 0 表示上坡，flg = 1 表示下坡
func validMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	i := 1
	for i < n && arr[i] > arr[i-1] {
		i++
	}
	// 如果上坡结束后 1 是水平；2 已经是最后一个元素（没有下坡）；3 没有上坡，都不是山峰数组
	if (i < n && arr[i] == arr[i-1]) || i == n || i == 1 {
		return false
	}
	for i < n && arr[i] < arr[i-1] {
		i++
	}
	return i == n
}
