package b_huawei_leetcode

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s) == 1 {
		return false
	}
	var arr []string
	for _, c := range s {
		str := string(c)
		if str == "(" {
			arr = append(arr, ")")
		} else if str == "[" {
			arr = append(arr, "]")
		} else if str == "{" {
			arr = append(arr, "}")
		} else if len(arr) == 0 || arrayPop(&arr) != str {
			return false
		}
	}
	return len(arr) == 0
}

func arrayPop(arr *[]string) string {
	if len(*arr) == 0 {
		return ""
	}
	x := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return x
}
