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

func checkMatch(arr [][]string, i, j int, match string) bool {
	for i := 0; i < len(match); i++ {
	}
	arr[i][j]
}

func solvep1(r io.Reader, match string) int {
	arr2d := inputs.Scan2DInput[string](r)
	lineLen := len(arr2d[0])
	// Add padding to inputs pkg
	for i := range arr2d {
		arr2d[i] = append(arr2d[i], repeatStrArr(padStr, len(match))...)
		arr2d[i] = append(repeatStrArr(padStr, len(match)), arr2d[i]...)
	}
	arr2d = append([][]string{{strings.Repeat(padStr, lineLen)}}, arr2d...)
	arr2d = append(arr2d, repeatStrArr(padStr, lineLen))

	for i := 0; i < len(arr2d); i++ {
		for j := 0; j < len(arr2d[i]); j++ {
			if arr2d[i][j] == "." {
				continue
			}

		}
	}

	return 0
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

		// spew.Dump(solvep1(bytes.NewReader(input)))
	})
}
