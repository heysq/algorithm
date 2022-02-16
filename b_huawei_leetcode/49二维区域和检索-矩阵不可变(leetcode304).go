package b_huawei_leetcode

type NumMatrix struct {
	matrix [][]int
}

func ConstructorV3(matrix [][]int) NumMatrix {

	var preMatrix = make([][]int, len(matrix))
	for i, row := range matrix {
		preMatrix[i] = make([]int, len(row)+1)
		for j, val := range row {
			preMatrix[i][j+1] = preMatrix[i][j] + val
		}
	}
	return NumMatrix{matrix: preMatrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	var sum int
	for i := row1; i <= row2; i++ {
		sum += this.matrix[i][col2+1] - this.matrix[i][col1]
	}
	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
