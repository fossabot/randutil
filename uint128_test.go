package randutil_test

import (
	"testing"

	"github.com/hnakamur/randutil"
)

func TestUint128ToZeroPaddedBase36(t *testing.T) {
	testCases := []struct {
		v    [16]byte
		want string
	}{
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			"0000000000000000000000000",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			"0000000000000000000000001",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 35},
			"000000000000000000000000z",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 36},
			"0000000000000000000000010",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 15},
			"00000000000000000000000zz",
		},
		{
			[16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, 16},
			"0000000000000000000000100",
		},
		{
			[16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 254},
			"f5lxx1zz5pnorynqglhzmsp32",
		},
		{
			[16]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255},
			"f5lxx1zz5pnorynqglhzmsp33",
		},
	}
	for _, tc := range testCases {
		got := randutil.Uint128ToZeroPaddedBase36(tc.v)
		if got != tc.want {
			t.Errorf("got %s, want %s", got, tc.want)
		}
	}
}
