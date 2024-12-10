package d10

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input []byte

/*

struct node {
	visited bool
	score int
}

checkedTrailHeads = [][]int
arr := make2darray()
arr = padArray(arr)
dirs = []int{
	0,1 R
	0,-1 L
	1,0 D
	0,-1 U
}

fn trailHeuristicsOk(arr []node, x, y, dx, dy int){
	want := []int{1,2,3,4,5,6,7,8,9}
	for i := range want
		for d := range dirs
			nx, ny := x+i*dx, y+i*dy
			if x, y == dx, dy: continue
			if arr[nx][ny] != want[i]
				return false
	return true
}

for i := range arr
	for j := range arr[i]
		if arr[i][j] == '0' && checkedTrailHeads[i][j] == nil
			checkedTrailHeads = [i][j]
			trailOk := trailHeuristicsOk(arr, i, j, dirs)
			if trailOk
				sum++
*/

// func solvep1(r io.Reader) int {
// 	b, err := io.ReadAll(r)
// 	if err != nil {
// 		panic(err)
// 	}
// 	_ = inputs.ParseGrid(string(b))
// 	return 0
// }

func TestSolve(t *testing.T) {
	t.Skip()
	t.Run("mock", func(t *testing.T) {
		in := `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

		got := solvep1(strings.NewReader(in))
		want := 36
		assert.Equal(t, want, got)
	})
}
