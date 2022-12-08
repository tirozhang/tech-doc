package _200_Number_of_Islands

import (
	"fmt"
)

var visited = make(map[string]int)

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]byte) (byte, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

// bfs实现
func numIslands(grid [][]byte) int {
	visited = map[string]int{}
	isLand := 0
	row, col := len(grid), len(grid[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			key := fmt.Sprintf("%d_%d", i, j)
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = 1
			if grid[i][j] == 1 {
				bfs(grid, i, j)
				isLand++
			}

		}
	}
	return isLand
}

func bfs(grid [][]byte, i, j int) {
	queue := []point{{i, j}}
	for len(queue) > 0 {
		land := queue[0]
		queue = queue[1:]
		for _, next := range dirs {
			isNextLand := next.add(land)
			if isWater, ok := isNextLand.at(grid); isWater == 0 || !ok {
				continue
			}
			key := fmt.Sprintf("%d_%d", isNextLand.i, isNextLand.j)
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = 1
			queue = append(queue, isNextLand)
		}

	}

}

// 递归实现宽度优先
func numIslands2(grid [][]byte) int {
	result, m, n := 0, len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				result++
				setIsLand(grid, i, j)
			}

		}
	}
	return result
}

func setIsLand(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] == 1 {
		grid[i][j] = 0
		setIsLand(grid, i, j-1)
		setIsLand(grid, i, j+1)
		setIsLand(grid, i-1, j)
		setIsLand(grid, i+1, j)
	}
}
