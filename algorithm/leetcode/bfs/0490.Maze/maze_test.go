package _490_Maze

import "testing"

func Test_NumIslands(t *testing.T) {
	t.Log(hasPath([][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}, []int{0, 4}, []int{3, 2})) //true
	t.Log(hasPath([][]int{
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{1, 1, 0, 1, 1},
		{0, 0, 0, 0, 0},
	}, []int{0, 4}, []int{1, 2})) //true
	t.Log(hasPath([][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}, []int{0, 0}, []int{1, 2})) //false
}
