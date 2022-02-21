package leetcode

func checkInclusion(s1 string, s2 string) bool {
	s1Map := make(map[byte]int)
	s2Map := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		s1Map[s1[i]]++
	}
	var left, right, valid int
	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := s1Map[c]; ok {
			s2Map[c]++
			if s2Map[c] == s1Map[c] {
				valid++
			}
		}
		for right-left >= len(s1) {
			if valid == len(s1Map) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := s1Map[d]; ok {
				if s2Map[d] == s1Map[d] {
					valid--
				}
				s2Map[d]--
			}
		}
	}
	return false
}
