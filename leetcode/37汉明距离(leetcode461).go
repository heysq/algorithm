package leetcode

func hammingDistance(x int, y int) int {
	var distance int
	for x := x ^ y; x != 0; x = x & (x - 1) {
		distance++
	}
	return distance
}
