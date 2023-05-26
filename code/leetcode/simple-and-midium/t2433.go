package simple_and_midium


func findArray(pref []int) []int {
	if len(pref) == 0 {
		return nil
	}
	ans := make([]int, len(pref))
	ans[0] = pref[0]
	ansXor := ans[0]
	for i:=1; i<len(ans); i++ {
		ans[i] = ansXor ^ pref[i]
		ansXor ^= ans[i]
	}
	return ans
}