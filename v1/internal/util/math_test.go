package util_test

import (
	"fmt"
	"testing"

	"github.com/foxcapades/gVersion/v1/internal/util"
)

func TestPowU8(t *testing.T) {
	tests := [...]struct {
		inp  uint8
		pow  uint8
		out  uint8
	}{
		{2, 0, 1},
		{2, 1, 2},
		{2, 2, 4},
		{2, 3, 8},
		{2, 4, 16},
		{2, 5, 32},
		{2, 6, 64},
		{2, 7, 128},

		{10, 1,10},
		{10, 2,100},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("(%d, %d) -> %d", test.inp, test.pow, test.out), func(t *testing.T) {
			val := util.PowU8(test.inp, test.pow)
			if val != test.out {
				t.Errorf("Expected %d^%d to equal %d, got %d", test.inp, test.pow, test.out, val)
			}
		})
	}
}
