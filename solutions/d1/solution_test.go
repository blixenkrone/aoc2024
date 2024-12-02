package d1

import (
	"bufio"
	"bytes"
	_ "embed"
	"slices"
	"testing"

	"blixenkrone/aoc2024/inputs"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input []byte

func prepareInput(in []byte) [][]int {

	sc := bufio.NewScanner(bytes.NewReader(in))
	sc.Split(bufio.ScanWords)
	k := 0
	pairs := make([][]int, 2)
	for sc.Scan() {
		out := sc.Text()
		intv := inputs.MustAtoi(out)
		pairs[k] = append(pairs[k], intv)

		if k > 0 {
			k--
			continue
		}
		k++
	}

	return pairs
}

func solveP2(in []byte) int {
	pairs := prepareInput(in)
	rightArrM := make(map[int]int)
	for _, num := range pairs[1] {
		rightArrM[num]++
	}

	var sum int
	for _, num := range pairs[0] {
		if v, ok := rightArrM[num]; ok {
			sum += (num * v)
		}
	}
	return sum
}

func solveP1(in []byte) int {
	pairs := prepareInput(in)

	for _, p := range pairs {
		slices.Sort(p)
	}

	k := 0
	var sum int
	for i := 0; i < len(pairs[0]); i++ {
		calc := inputs.Abs(pairs[0][k] - pairs[1][k])
		sum += calc
		k++
	}

	return sum
}

func TestSolveP2(t *testing.T) {

	t.Run("mock", func(t *testing.T) {
		in := []byte(`3   4
4   3
2   5
1   3
3   9
3   3`)

		got := solveP2(in)
		assert.Equal(t, got, 31)

	})
	t.Run("vreal", func(t *testing.T) {
		spew.Dump(solveP2(input))

	})
}

func TestSolveP1(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		in := []byte(`3   4
4   3
2   5
1   3
3   9
3   3`)

		got := solveP1(in)
		assert.Equal(t, got, 11)

	})
	t.Run("real", func(t *testing.T) {

		out := solveP1(input)
		spew.Dump(out)
	})
}
