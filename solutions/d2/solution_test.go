package d2

import (
	"bytes"
	_ "embed"
	"io"
	"slices"
	"strings"
	"testing"

	"blixenkrone/aoc2024/inputs"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input []byte

func prepareInput(r io.Reader) [][]int {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	var out [][]int
	s1 := strings.Split(string(b), "\n")
	for _, s := range s1 {
		if s == "" {
			continue
		}
		var intarr []int
		s2 := strings.Split(s, " ")
		for _, v := range s2 {
			intarr = append(intarr, inputs.MustAtoi(v))
		}
		out = append(out, intarr)
	}
	return out
}

func checkAscDesc(arr []int) bool {
	isDesc := arr[0] > arr[1]

	if arr[0] == arr[1] {
		return false
	}

	for i := 0; i < len(arr)-1; i++ {
		diff := inputs.Abs(arr[i] - arr[i+1])
		if diff > 3 || diff <= 0 {
			return false
		}
		if isDesc && arr[i] > arr[i+1] {
			continue
		}
		if !isDesc && arr[i] < arr[i+1] {
			continue
		}

		return false
	}
	return true

}

func solve(r io.Reader) int {
	int2dArr := prepareInput(r)
	sum := 0
	for _, intArr := range int2dArr {
		if checkAscDesc(intArr) {
			sum++
		}

	}
	return sum
}

func solvep2(r io.Reader) int {
	int2dArr := prepareInput(r)
	sum := 0
	for _, intArr := range int2dArr {
		if checkAscDesc(intArr) {
			sum++
			continue
		}
		for i := 0; i < len(intArr); i++ {
			perm := slices.Clone(intArr)
			perm = slices.Delete(perm, i, i+1)
			if checkAscDesc(perm) {
				sum++
				break
			}
		}

	}
	return sum
}

func TestSolveP2(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		in := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

		got := solvep2(strings.NewReader(in))
		assert.Equal(t, 4, got)

	})
	t.Run("real", func(t *testing.T) {

		spew.Dump(solvep2(bytes.NewReader(input)))
	})
}

func TestSolveP1(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		in := `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

		got := solve(strings.NewReader(in))
		assert.Equal(t, 2, got)

	})
	t.Run("real", func(t *testing.T) {
		spew.Dump(solve(bytes.NewReader(input)))

	})
}
