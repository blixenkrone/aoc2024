package d9

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

func prepareInput(r io.Reader) string {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func translateFileBlocksToArray(s string) []int {
	var result []int
	idx := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
			continue
		}
		strv := inputs.MustAtoi(string(s[i]))

		if i%2 == 0 { // File block
			// Append `idx` (the current block index) `strv` times
			for j := 0; j < strv; j++ {
				result = append(result, idx)
			}
			idx++
			continue
		}
		// Free space
		// Append `-1` (representing free space) `strv` times
		for j := 0; j < strv; j++ {
			result = append(result, -1)
		}
	}

	return result
}

func allocateFreeSpaceArray(arr []int) []int {
	left := 0
	right := len(arr) - 1

	for left < right {
		for left < len(arr) && arr[left] != -1 {
			left++
		}
		for right >= 0 && arr[right] == -1 {
			right--
		}

		// Swap the free space and the block if pointers are valid
		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	return arr
}

func countIdxValuesArrayP1(arr []int) int {
	sum := 0
	for i, val := range arr {
		if val != -1 { // Ignore free space
			sum += val * i
		}
	}
	return sum
}

func solvep1(r io.Reader) int {
	// Step 1: Prepare input
	input := prepareInput(r)

	// Step 2: Translate the file blocks into an array
	blocks := translateFileBlocksToArray(input)

	// Step 3: Allocate free space
	allocated := allocateFreeSpaceArray(blocks)

	// Step 4: Calculate the checksum
	return countIdxValuesArrayP1(allocated)
}

func TestSolveP1(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		in := strings.NewReader(`2333133121414131402`)

		got := solvep1(in)
		assert.Equal(t, got, 2858)
	})
	t.Run("real", func(t *testing.T) {
		// got := solve(bytes.NewReader(input))
		// fasdmt.Println(got)
	})
}
