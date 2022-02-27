package leetcode

func findCelebrity(n int) int {
	var cand int
	for other := 1; other < n; other++ {
		if knows(cand, other) || !knows(other, cand) {
			cand = other
		}
	}

	// cand 是候选人
	for other := 0; other < n; other++ {
		if cand == other {
			continue
		}

		if knows(cand, other) || !knows(other, cand) {
			return -1
		}
	}
	return cand
}

func knows(n, m int) bool {
	return false
}
