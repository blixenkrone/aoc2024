package d9

import (
	"blixenkrone/aoc2024/inputs"
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func translateFileBlocksToArrayP2(s string) [][]int {
	var result [][]int
	idx := 0

	for i := 0; i < len(s); i++ {
		var item []int
		if s[i] == '\n' {
			continue
		}
		strv := inputs.MustAtoi(string(s[i]))

		if i%2 == 0 { // File block
			// Append `idx` (the current block index) `strv` times
			for j := 0; j < strv; j++ {
				item = append(item, idx)
			}
			result = append(result, item)
			idx++
			continue
		}
		// Free space
		// Append `-1` (representing free space) `strv` times
		for i := 0; i < strv; i++ {
			item = append(item, -1)
		}
		result = append(result, item)
	}

	return result
}

func allocateFreeSpaceArrayP2(arr [][]int) []int {
	// k := len(arr) - 1
	moves := len(arr) - 1
	idx := 0
	for moves > 0 {
		if len(arr[idx]) <= 0 {
			continue
		}
		if arr[idx][0] < 0 {
			for j := len(arr); j >= 0; j-- {
				if arr[j][0] < 0 {
					continue
				}
				if len(arr[j]) > len(arr[idx]) {

				}
			}
		}
	}
	return nil
}

func countIdxValuesArrayP2(arr []int) int {
	sum := 0
	for i, val := range arr {
		if val != -1 { // Ignore free space
			sum += val * i
		}
	}
	return sum
}

func solvep2(r io.Reader) int {
	// Step 1: Prepare input
	input := prepareInput(r)

	// Step 2: Translate the file blocks into an array
	blocks := translateFileBlocksToArray(input)

	// Step 3: Allocate free space
	allocated := allocateFreeSpaceArray(blocks)

	// Step 4: Calculate the checksum
	return countIdxValuesArrayP1(allocated)
}
func allocateFreeSpaceWithMove(s string) string {
	n := len(s)
	arr := []byte(s)

	// Parse the input to identify files and their spans
	files := []struct {
		id    byte
		start int
		end   int
		size  int
	}{}

	// Identify files
	i := 0
	for i < n {
		if arr[i] != '.' {
			start := i
			for i < n && arr[i] == arr[start] {
				i++
			}
			files = append(files, struct {
				id    byte
				start int
				end   int
				size  int
			}{id: arr[start], start: start, end: i - 1, size: i - start})
		} else {
			i++
		}
	}

	// Sort files by ID in descending order
	sort.Slice(files, func(i, j int) bool {
		return files[i].id > files[j].id
	})

	// Attempt to move files
	for _, file := range files {
		fileSize := file.size

		// Find a suitable span of free space to the left of the file
		for j := 0; j < file.start; j++ {
			// Check if this span of free space is large enough
			isSpan := true
			for k := 0; k < fileSize; k++ {
				if j+k >= file.start || arr[j+k] != '.' {
					isSpan = false
					break
				}
			}

			if isSpan {
				// Move the file into this span
				for k := 0; k < fileSize; k++ {
					arr[j+k] = file.id
				}
				// Clear the original file's space
				for k := file.start; k <= file.end; k++ {
					arr[k] = '.'
				}
				break
			}
		}
	}

	return string(arr)
}
func TestSolveP2(t *testing.T) {
	t.Run("mock", func(t *testing.T) {

		in := strings.NewReader(`2333133121414131402`)

		got := solvep2(in)
		assert.Equal(t, got, 1928)
	})
	t.Run("real", func(t *testing.T) {

		got := solvep2(bytes.NewReader(input))
		fmt.Println(got)
	})
}
