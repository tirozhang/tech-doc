package _200_Number_of_Islands

import "testing"

func Test_NumIslands(t *testing.T) {
	t.Log(numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 0, 0, 1, 1},
		{0, 0, 0, 0, 1},
	})) // 2
	t.Log(numIslands([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	})) // 3
	t.Log(numIslands2([][]byte{
		{1, 1, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 1, 1},
	})) // 3
}
