package simple_and_midium

import (
	"bytes"
	"strings"
)

func convert(s string, numRows int) string {
	n := len(s)
	var ans strings.Builder
	ans.Grow(n)
	if numRows == 1 {
		return s
	} else if numRows == 2 {
		for i := 0; i < n; i += 2 {
			ans.WriteByte(s[i])
		}
		for i := 1; i < n; i += 2 {
			ans.WriteByte(s[i])
		}
		return ans.String()
	}

	for i := 0; i < numRows; i++ {
		switch i {
		case 0:
			for j := i; j < n; j += 2*numRows - 2 {
				ans.WriteByte(s[j])
			}
		case numRows-1:
			for j := i; j < n; j += 2*numRows - 2 {
				ans.WriteByte(s[j])
			}
		default:
			for j := i; j < n; j += 2*numRows - 2 {
				ans.WriteByte(s[j])
				adjChIdx := j + (numRows - i)
				if adjChIdx < n {
					ans.WriteByte(s[adjChIdx])
				}
			}
		}
	}
	return ans.String()
}

func convert2(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}