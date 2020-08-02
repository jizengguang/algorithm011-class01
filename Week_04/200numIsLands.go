package Week_04

func numIslands(grid [][]byte) int {
	var count int
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] == '0' {
				continue
			}
			count++
			dfsFind(grid, line, column)
		}
	}
	return count
}

func dfsFind(gird [][]byte, i, j int) {
	if i < 0 || i >= len(gird) || j < 0 || j >= len(gird[i]) {
		return
	}

	if gird[i][j] == '0' {
		return
	}

	gird[i][j] = '0'

	dfsFind(gird, i-1, j)
	dfsFind(gird, i+1, j)
	dfsFind(gird, i, j-1)
	dfsFind(gird, i, j+1)
}
