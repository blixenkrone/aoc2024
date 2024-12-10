package d10

import (
	"blixenkrone/aoc2024/inputs"
	"io"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

const (
	Day  = "10"
	Year = "2024"
)

const (
	north byte = 1 << iota
	south
	east
	west
)

type point struct{ i, j int }

func inBounds(i, j, iMax, jMax int) bool { return i > -1 && i < iMax && j > -1 && j < jMax }

func uniqueDFS(startI, startJ int, grid [][]byte, set map[point]struct{}) (score int) {
	var (
		iStack, jStack [10]int
		visited        [10]byte
		iMax           = len(grid)
		jMax           = len(grid[0])
	)
	iStack[0] = startI
	jStack[0] = startJ

	for n, cur := 0, byte('0'); n > -1; n, cur = n-1, cur-1 {
	next:
		for {
			switch {
			case visited[n]&north == 0:
				visited[n] |= north
				iStack[n+1] = iStack[n] - 1
				jStack[n+1] = jStack[n]
			case visited[n]&south == 0:
				visited[n] |= south
				iStack[n+1] = iStack[n] + 1
				jStack[n+1] = jStack[n]
			case visited[n]&east == 0:
				visited[n] |= east
				iStack[n+1] = iStack[n]
				jStack[n+1] = jStack[n] + 1
			case visited[n]&west == 0:
				visited[n] |= west
				iStack[n+1] = iStack[n]
				jStack[n+1] = jStack[n] - 1
			default:
				visited[n] = 0
				break next
			}

			if inBounds(iStack[n+1], jStack[n+1], iMax, jMax) && grid[iStack[n+1]][jStack[n+1]] == cur+1 {
				n++
				cur++
				if n == 9 {
					set[point{iStack[n], jStack[n]}] = struct{}{}
					break next
				}
			}
		}
	}
	return len(set)
}

func dfs(startI, startJ int, grid [][]byte) (score int) {
	var (
		iStack, jStack [10]int
		visited        [10]byte
		iMax           = len(grid)
		jMax           = len(grid[0])
	)
	iStack[0] = startI
	jStack[0] = startJ

	for n, cur := 0, byte('0'); n > -1; n, cur = n-1, cur-1 {
	next:
		for {
			switch {
			case visited[n]&north == 0:
				visited[n] |= north
				iStack[n+1] = iStack[n] - 1
				jStack[n+1] = jStack[n]
			case visited[n]&south == 0:
				visited[n] |= south
				iStack[n+1] = iStack[n] + 1
				jStack[n+1] = jStack[n]
			case visited[n]&east == 0:
				visited[n] |= east
				iStack[n+1] = iStack[n]
				jStack[n+1] = jStack[n] + 1
			case visited[n]&west == 0:
				visited[n] |= west
				iStack[n+1] = iStack[n]
				jStack[n+1] = jStack[n] - 1
			default:
				visited[n] = 0
				break next
			}

			if inBounds(iStack[n+1], jStack[n+1], iMax, jMax) && grid[iStack[n+1]][jStack[n+1]] == cur+1 {
				n++
				cur++
				if n == 9 {
					score++
					break next
				}
			}
		}
	}
	return score
}

func solvep1(r io.Reader) int {
	var count int
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	grid := inputs.ParseGrid(string(b))
	spew.Dump(grid)
	// set := make(map[point]struct{})
	// for i, row := range grid {
	// 	for j, c := range row {
	// 		if c == '0' {
	// 			clear(set)
	// 			score := uniqueDFS(i, j, grid, set)
	// 			count += score
	// 		}
	// 	}
	// }
	return count
}

func solvep2(r io.Reader) int {
	var count int
	// b, err := io.ReadAll(r)
	// if err != nil {
	// 	panic(err)
	// }
	// grid := inputs.ParseGrid(string(b))
	// for i, row := range grid {
	// 	for j, c := range row {
	// 		if c == '0' {
	// 			score := dfs(i, j, grid)
	// 			count += score
	// 		}
	// 	}
	// }
	return count
}

func TestSolveP1(t *testing.T) {

	t.Run("mock", func(t *testing.T) {
		in := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		got := solvep1(strings.NewReader(in))
		want := 36
		assert.Equal(t, want, got)
	})
}
