package leetcode

/*
异或一下，
然后计算1的个数
f(x) 表示 xx 和 x-1x−1 进行与运算所得的结果（即 f(x)=x~\&~(x-1)f(x)=x & (x−1)），那么 (x) 恰为 x 删去其二进制表示中最右侧的 1 的结果。
*/
func hammingDistance(x int, y int) int {
	var distance int
	for x := x ^ y; x != 0; x = x & (x - 1) {
		distance++
	}
	return distance
}
