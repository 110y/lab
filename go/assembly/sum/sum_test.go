package sum_test

import (
	"testing"

	"github.com/110y/lab/go/assembly/sum"
)

func TestSum(t *testing.T) {
	tests := map[string]struct {
		a        int64
		b        int64
		expected int64
	}{
		"1": {
			a:        1,
			b:        2,
			expected: 3,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			if want, got := test.expected, sum.Sum(test.a, test.b); want != got {
				t.Errorf("want %d, but got %d", want, got)
			}
		})
	}
}
