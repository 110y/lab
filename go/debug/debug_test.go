package debug_test

import (
	"testing"

	"github.com/110y/lab/go/debug"
)

func TestDouble(t *testing.T) {
	tests := map[string]struct {
		n        int64
		expected int64
	}{
		"1": {
			n:        2,
			expected: 4,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			if want, got := test.expected, debug.Double(test.n); want != got {
				t.Errorf("want %d, but goot %d", want, got)
			}
		})
	}
}
