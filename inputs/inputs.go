package inputs

import (
	"io"
	"strconv"
	"strings"
)

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

func Scan2DInput[T ~string](in io.Reader) [][]T {
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
