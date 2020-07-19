package Week_04

func numIslands(grid [][]byte) int {
	var count int = 0
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

func dfsFind(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfsFind(grid, i-1, j)
	dfsFind(grid, i+1, j)
	dfsFind(grid, i, j-1)
	dfsFind(grid, i, j+1)

}
