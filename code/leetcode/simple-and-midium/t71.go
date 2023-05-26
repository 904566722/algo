package simple_and_midium

//  // -> /
//  ./ -> 直接把刚入栈的 . 去除
//  ../ -> 回退到上一个 /
// .../ -> 视作合法
func simplifyPath(path string) string {
	var st []byte
	// 因为这个方法都是以 '/' 来判断往前回退，所以在最后补充一个 '/'
	// 保证字符串以 '/' 结尾
	path += "/"
	for i := range path {
		if len(st) == 0 || path[i] != '/' {
			st = append(st, path[i])
		} else {
			// len(st) > 0
			// path[i] == '/'

			top := len(st) - 1
			sec := top - 1
			thr := sec - 1

			if st[top] == '/' {
				// //
				// nothing to do
				// dont push this '/'
			} else if thr >= 0 && thr == '.' {
				// .../
			} else if sec >= 0 && st[sec] == '.' {
				// ../
				// st = st[:len(st)-3]
				for st[len(st)-1] != '/' {
					st = st[:len(st)-1]
				}
				// 如果此时 len(st) == 1
				// 说明已经是根目录 / 无法继续往上找目录，到此结束

				// 如果此时 len(st) > 1
				// 也就说明前面还有目录     xxxxa/
				// 需要继续往前找到一个 '/'
				if len(st) > 1 {
					st = st[:len(st)-1]
					for st[len(st)-1] != '/' {
						st = st[:len(st)-1]
					}
				}
			} else if st[top] == '.' {
				// ./
				st = st[:len(st)-1]
			} else {
				st = append(st, path[i])
			}
		}
	}
	if len(st) != 1 &&  st[len(st)-1] == '/' {
		return string(st[:len(st)-1])
	}
	return string(st)
}