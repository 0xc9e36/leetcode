/**
 * @Author: wei.tan
 * @Description:
 * @File:  solution
 * @Version: 1.0.0
 * @Date: 2019-10-21 21:12
 */

package MagicSquaresInGrid

func numMagicSquaresInside(grid [][]int) int {
	line, column := len(grid), len(grid[0])

	if line < 3 || column < 3 {
		return 0
	}

	res := 0
	for i := 0; i < line-2; i++ {
		for j := 0; j < column-2; j++ {
			if IsMagicSquare(grid, i, j) {
				res++
			}
		}
	}

	return res
}

func IsMagicSquare(grid [][]int, i, j int) bool {
	used := [10]bool{}

	val := 0
	for dx := 0; dx < 3; dx++ {
		lineval, columnval := 0, 0
		for dy := 0; dy < 3; dy++ {
			val = grid[i+dx][j+dy]
			if val < 1 || val > 9 {
				return false
			}
			used[val] = true
			lineval += grid[i+dx][j+dy]
			columnval += grid[i+dy][j+dx]
		}

		if lineval != 15 || columnval != 15 {
			return false
		}
	}

	for i := 1; i <= 9; i++ {
		if !used[i] {
			return false
		}
	}

	if 15 != grid[i][j]+grid[i+1][j+1]+grid[i+2][j+2] || 15 != grid[i][j+2]+grid[i+1][j+1]+grid[i+2][j] {
		return false
	}

	return true

}
