package _490_Maze

import (
	"fmt"
	"leetcode/help_func"
)

type Point = help_func.Point

var visited = make(map[string]int)

func hasPath(maze [][]int, start []int, destination []int) bool {
	var key string
	visited = map[string]int{}
	queue := []Point{{start[0], start[1]}}

	for len(queue) > 0 {
		topStep := queue[0]
		queue = queue[1:]

		for k, next := range help_func.Dirs {
			nowStep := topStep
			for {
				nextStep := nowStep.Add(next)
				if isWall, ok := nextStep.At(maze); !ok || isWall == 1 {
					break
				}
				nowStep = nextStep
			}
			if nowStep.X == destination[0] && nowStep.Y == destination[1] {
				return true
			}
			key = fmt.Sprintf("%d_%d_%d", nowStep.X, nowStep.Y, k%2)
			if _, ok := visited[key]; ok {
				continue
			}
			visited[key] = 1
			queue = append(queue, nowStep)

		}
	}
	return false
}
