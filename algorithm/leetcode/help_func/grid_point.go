package help_func

type Point struct {
	X, Y int
}

var Dirs = [4]Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p *Point) Add(r Point) Point {
	return Point{p.X + r.X, p.Y + r.Y}
}

func (p *Point) At(grid [][]int) (int, bool) {
	if p.X < 0 || p.X >= len(grid) {
		return 0, false
	}

	if p.Y < 0 || p.Y >= len(grid[p.X]) {
		return 0, false
	}

	return grid[p.X][p.Y], true
}
