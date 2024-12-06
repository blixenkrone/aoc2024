package d6

import (
	"blixenkrone/aoc2024/inputs"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input []byte

func prepareInput(r io.Reader) [][]string {
	grid := inputs.PadGrid(inputs.Scan2DInput[string](r), "?", 1)
	return grid

}

type direction int

const (
	_ direction = iota - 1
	up
	down
	right
	left
)

/*
// order matters
dirs = [
	-1,0,   up
	0,1, right
	-1,0, down
	0,-1, left
]
grid := []string
px,py := 2,1
dirIdx = ?
dir = dirs[dirIdx]
curr = grid[px, py]

for i := 1; ; ;
	nx, ny = peekNext(nx, ny, dir, i)
	if grid[nx][ny] == "#"
		dir, dirIdx = nextDir(dir, dirIdx)
		i = 1
		continue
	if grid[nx][ny] == "."
		countField()
		move(nx, ny, dir)
		curr = "."
		continue
	if grid[nx, ny] == "?"
		break
*/

var dirs = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func initVars(grid [][]string) (int, int, direction) {
	dirM := map[string]direction{
		"^": up,
		">": right,
		"v": down,
		"<": left,
	}
	for i, v1 := range grid {
		for j, v2 := range v1 {
			if v, ok := dirM[v2]; ok {
				return i, j, v
			}
		}
	}

	panic("init vars failed")

}

func peekNext(x, y, dx, dy, i int) (int, int) {
	nx, ny := x+(i*dx), y+(i*dy)
	return nx, ny
}

func nextDir(dir direction) direction {
	nextDir := dir + 1
	if int(dir) == len(dirs)-1 {
		nextDir = 0
	}
	return nextDir
}

func solve(r io.Reader) int {
	grid := prepareInput(r)
	px, py, dir := initVars(grid)
	curr := grid[px][py]
	fmt.Printf("initial curr %d, %d \n", px, py)
	// dir := [][2]int{}

	// distinctTilesM := make(map[int]struct{})

	var sum int
	for {
		nx, ny := peekNext(px, py, dirs[dir][0], dirs[dir][1], 1)
		fmt.Printf("peeked %d,%d i j \n", nx, ny)
		curr = grid[nx][ny]
		if curr == "#" {
			dir = nextDir(dir)
			spew.Dump("new dir " + strconv.Itoa(int(dir)))
			continue
		}
		if curr == "." || curr == "X" {
			// count field
			// move()
			if curr == "." {
				sum++
			}
			grid[px][py] = "X"
			curr = grid[nx][ny]
			px, py = nx, ny
			fmt.Printf("new pos %d,%d i j \n", nx, ny)
			fmt.Println("curr " + curr)
			continue
		}
		if curr == "?" {
			fmt.Printf("DONE at %d,%d\n", nx, ny)
			sum += 1
			printGrid(grid)
			break
		}
	}
	return sum
}

func printGrid(grid [][]string) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			fmt.Print(grid[row][column], " ")
		}
		fmt.Print("\n")
	}

}

func TestSolveP1(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		// remember theres padding
		// t.Skip()
		in := strings.NewReader(`
....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`)

		got := solve(in)
		assert.Equal(t, got, 41)
	})
	t.Run("mock", func(t *testing.T) {
		t.Skip()
		// remember theres padding
		got := solve(bytes.NewReader(input))
		spew.Dump(got)
		// assert.Equal(t, got, 41)
	})

}
