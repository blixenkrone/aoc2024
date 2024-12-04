package inputs

import (
	"bufio"
	"io"
	"strconv"
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

func Scan2DInput[T ~string](in io.Reader) [][]T {
	sc := bufio.NewScanner(in)
	sc.Split(bufio.ScanRunes)
	var out [][]T
	var row []T
	for sc.Scan() {
		str := sc.Text()
		if str == "\t" || str == " " || str == "" {
			continue
		}
		if str == "\n" {
			out = append(out, row)
			row = nil
			continue
		}
		row = append(row, T(str))
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	return out

}

// func PrepareInput[T comparable](in io.Reader) T {
// 	sc := bufio.NewScanner(in)
// 	sc.Split(bufio.ScanWords)
// 	k := 0
// 	pairs := make([][]int, 2)
// 	for sc.Scan() {
// 		out := sc.Text()
// 		intv := inputs.MustAtoi(out)
// 		pairs[k] = append(pairs[k], intv)

// 		if k > 0 {
// 			k--
// 			continue
// 		}
// 		k++
// 	}
// 	return pairs
// }
