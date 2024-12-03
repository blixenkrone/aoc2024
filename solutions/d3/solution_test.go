package d2

import (
	"blixenkrone/aoc2024/inputs"
	"bytes"
	_ "embed"
	"io"
	"regexp"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

//go:embed input.txt
var input []byte

var (
	findMulRegexP1 = regexp.MustCompile(`mul\(\d+\,\d+\)`)
	findMulRegexP2 = regexp.MustCompile(`(?:mul\(\d+,\d+\))|(?:do\(\)|don\'t\(\))`)
)

func solve(r io.Reader) int {
	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	mulStrArr := findMulRegexP1.FindAllString(string(b), -1)
	var sum int

	for _, m := range mulStrArr {
		m = strings.TrimPrefix(m, "mul(")
		m = strings.TrimSuffix(m, ")")
		intv := strings.SplitAfter(m, ",")
		if len(intv) > 2 {
			panic("intv is too big " + strings.Join(intv, ", "))
		}
		a, b := inputs.MustAtoi(strings.TrimSuffix(intv[0], ",")), inputs.MustAtoi(intv[1])
		sum += a * b
	}
	return sum
}

func solvep2(r io.Reader) int {

	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	mulStrArr := findMulRegexP2.FindAllString(string(b), -1)
	spew.Dump(mulStrArr)
	active := true
	var sum int

	calcFn := func(str string) int {
		str = strings.TrimPrefix(str, "mul(")
		str = strings.TrimSuffix(str, ")")
		intv := strings.SplitAfter(str, ",")
		if len(intv) > 2 {
			panic("intv is too big " + strings.Join(intv, ", "))
		}
		a, b := inputs.MustAtoi(strings.TrimSuffix(intv[0], ",")), inputs.MustAtoi(intv[1])
		return a * b
	}
	for _, m := range mulStrArr {
		if m == "do()" {
			active = true
			continue
		}
		if m == `don't()` {
			active = false
			continue
		}

		if strings.HasPrefix(m, "mul") && active {
			sum += calcFn(m)
			continue
		}
	}

	return sum
}

func TestSolveP2(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		// in := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
		// got := solvep2(strings.NewReader(in))
		// assert.Equal(t, 48, got)

	})
	t.Run("mock2", func(t *testing.T) {
		// in := `+&don't()select()^)}$!who()+>mul(481,439)- *mul(332,923)when()-why()!mul(625,335)?mul(156,134)why()*'$:mul(104,186)^?#>/mul(605,863),,?mul(161,457)!#~,#~>/do()mul(514how()where(),'$;'$+)mul(207,90)&`
		// got := solvep2(strings.NewReader(in))
		// assert.Equal(t, 48, got)

	})
	t.Run("real", func(t *testing.T) {
		spew.Dump(solvep2(bytes.NewReader(input)))
	})
}
func TestSolveP1(t *testing.T) {
	t.Run("mock", func(t *testing.T) {
		in := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
		got := solve(strings.NewReader(in))
		assert.Equal(t, 161, got)

	})
	t.Run("real", func(t *testing.T) {
		spew.Dump(solve(bytes.NewReader(input)))
	})
}
