package day11

import "testing"

func TestCalculate(t *testing.T) {
	testCases := []struct {
		line       string
		iterations int
		want       int
	}{
		{
			line:       "125 17",
			iterations: 25,
			want:       55312,
		},
		{
			line:       "5910927 0 1 47 261223 94788 545 7771",
			iterations: 25,
			want:       193607,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.line, func(t *testing.T) {
			res := CalcWithMemo(tt.line, tt.iterations)

			if res != tt.want {
				t.Errorf("Got %d, want %d", res, tt.want)
			}
		})
	}
}
