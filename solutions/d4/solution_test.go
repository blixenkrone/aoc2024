package d4

import (
	"blixenkrone/aoc2024/inputs"
	_ "embed"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input []byte

const (
	padStr = "."
)

func repeatStrArr(s string, l int) (out []string) {
	for i := 0; i < l; i++ {
		out = append(out, s)
	}
	return
}

// Add n padding around a grid with the specified character
func padGrid(grid [][]string, n int) [][]string {
	// Add padding to inputs pkg
	for i := range grid {
		grid[i] = append(grid[i], repeatStrArr(padStr, n)...)
		grid[i] = append(repeatStrArr(padStr, n), grid[i]...)
	}
	lineLen := len(grid[0])
	grid = append([][]string{repeatStrArr(padStr, lineLen)}, grid...)
	grid = append(grid, repeatStrArr(padStr, lineLen))
	return grid
}

// Check all 8 directions iteratively
var directions = [][2]int{
	{0, 1},   // Right
	{0, -1},  // Left
	{1, 0},   // Down
	{-1, 0},  // Up
	{1, 1},   // Diagonal down-right
	{1, -1},  // Diagonal down-left
	{-1, 1},  // Diagonal up-right
	{-1, -1}, // Diagonal up-left
}

func solvep1(r io.Reader, match string) int {
	arr2d := inputs.Scan2DInput[string](r)
	paddedArr := padGrid(arr2d, len(match))
	sum := countWordOccurrences(paddedArr, match)
	return sum
}

// Check if a word exists in the grid starting from (x, y) in a given direction
func checkDirection(grid [][]string, x, y int, dx, dy int, word string) bool {
	for i, wchar := range word {
		// Calc angles/direction according to word[idx] char
		nx, ny := x+i*dx, y+i*dy
		wstr := string(wchar)
		// Out of bounds or mismatched character
		// if nx < 0 || ny < 0 || nx >= len(grid) || ny >= len(grid[0]) || grid[nx][ny] != str {
		if grid[nx][ny] != wstr {
			return false
		}
	}
	return true
}

// Count all occurrences of the word in the grid
func countWordOccurrences(grid [][]string, word string) int {
	count := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if grid[i][j] != string(word[0]) && grid[i][j] != string(word[len(word)-1]) {
					continue
				}
				if checkDirection(grid, i, j, dir[0], dir[1], word) {
					count++
				}
			}
		}
	}
	return count
}

func TestSolveP2(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		in := `.M.S......
..A..MSMS.
.M.S.MAA..
..A.ASMSM.
.M.S.M....
..........
S.S.S.S.S.
.A.A.A.A..
M.M.M.M.M.
..........`

		got := solvep1(strings.NewReader(in), "XMAS")
		assert.Equal(t, 18, got)

	})
	t.Run("real", func(t *testing.T) {
		// spew.Dump(solvep1(bytes.NewReader(input), "XMAS"))
	})
}
func TestSolveP1(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		in := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

		got := solvep1(strings.NewReader(in), "XMAS")
		assert.Equal(t, 18, got)

	})
	t.Run("real", func(t *testing.T) {
		// spew.Dump(solvep1(bytes.NewReader(input), "XMAS"))
	})
}
