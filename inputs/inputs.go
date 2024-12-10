package inputs

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// func MustUnicode[T ~int](i T) string {
// 	r := rune(i)
// 	return string(r)
// }

func Abs[T ~int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func MustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Reverse[T ~string](in T) T {
	runes := []rune(in)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return T(runes)
}

// Add n padding around a grid with the specified character
func PadGrid(grid [][]string, padChar string, n int) [][]string {
	repeatStrArr := func(s string, l int) (out []string) {
		for i := 0; i < l; i++ {
			out = append(out, s)
		}
		return
	}
	// Add padding to inputs pkg
	for i := range grid {
		grid[i] = append(grid[i], repeatStrArr(padChar, n)...)
		grid[i] = append(repeatStrArr(padChar, n), grid[i]...)
	}
	lineLen := len(grid[0])
	grid = append([][]string{repeatStrArr(padChar, lineLen)}, grid...)
	grid = append(grid, repeatStrArr(padChar, lineLen))
	return grid
}

func ParseGrid(input string) [][]int {
	// Split the input into lines
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]int, len(lines))

	for i, line := range lines {
		// Split each line into individual characters
		grid[i] = make([]int, len(line))
		for j, char := range line {
			// Convert each character to an integer
			num, err := strconv.Atoi(string(char))
			if err != nil {
				panic(fmt.Sprintf("Invalid character in input: %s", string(char)))
			}
			grid[i][j] = num
		}
	}

	return grid
}
func Scan2DInput[T ~string | ~int](in io.Reader) [][]T {
	b, err := io.ReadAll(in)
	if err != nil {
		panic(err)
	}
	// Split the string by newline characters to separate rows
	rows := strings.Split(strings.TrimSpace(string(b)), "\n")
	// Create a 2D slice of runes
	matrix := make([][]T, len(rows))
	for i, row := range rows {
		row = strings.TrimSpace(row)
		// v := T(row)
		matrix[i] = make([]T, len(row))
		for j, ch := range row {
			matrix[i][j] = T(ch)
		}
		// matrix[i] = append(matrix[i], v) // Convert each row to a slice
	}
	return matrix
}
