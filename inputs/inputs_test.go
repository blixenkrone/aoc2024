package inputs

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan2DInput(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		in := `A
		B`

		got := Scan2DInput[string](strings.NewReader(in))
		assert.Equal(t, [][]string{{"A"}, {"B"}}, got)
	})
	t.Run("2", func(t *testing.T) {
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

		got := Scan2DInput[string](strings.NewReader(in))
		assert.Equal(t, len(got), 10)
		assert.Equal(t, len(got[0]), 10)
	})

}
