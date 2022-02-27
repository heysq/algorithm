package leetcode

func equationsPossible(equations []string) bool {
	uf := NewUF(26)
	for _, i := range equations {
		if i[1] == '=' {
			uf.Union(int(i[0]-'a'), int(i[3]-'a'))
		}
	}

	for _, i := range equations {
		if i[1] == '!' {
			if uf.Connected(int(i[0]-'a'), int(i[3]-'a')) {
				return false
			}
		}
	}
	return true
}
