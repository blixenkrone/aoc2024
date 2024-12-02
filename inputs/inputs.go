package inputs

import (
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
