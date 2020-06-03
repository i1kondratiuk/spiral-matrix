package spiral

import (
	"reflect"
	"testing"
)

var TestData = []struct {
	in  [][]int
	out []int
}{
	{
		in: [][]int{
			{1, 2},
			{3, 4},
		},
		out: []int{1, 2, 4, 3},
	},
	{
		in: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
		out: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
	},
	{
		in: [][]int{
			{1, 2, 3, 1},
			{4, 5, 6, 4},
			{7, 8, 9, 7},
			{7, 8, 9, 7},
		},
		out: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
	},
	{
		in: [][]int{
			{1, 2, 3, 1, 0},
			{4, 5, 6, 4, 0},
			{7, 8, 9, 7, 0},
			{7, 8, 9, 7, 0},
			{7, 8, 9, 7, 0},
		},
		out: []int{1, 2, 3, 1, 0, 0, 0, 0, 0, 7, 9, 8, 7, 7, 7, 4, 5, 6, 4, 7, 7, 9, 8, 8, 9},
	},
}

func TestSpiralOneDim(t *testing.T) {
	for _, test := range TestData {
		res, _ := Spiral(test.in)
		AssertEqual(t, test.out, res)
	}
}

func AssertEqual(t *testing.T, a interface{}, b interface{}) {
	if reflect.DeepEqual(a, b) {
		return
	}

	t.Errorf("Expected %v (type %v), received %v (type %v)", a, reflect.TypeOf(a), b, reflect.TypeOf(b))
}
