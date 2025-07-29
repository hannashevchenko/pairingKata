package categorizer

import (
	"testing"
)

func TestCategorizer(t *testing.T) {
	tests := map[string]struct {
		input int
		want  Category
	}{
		"input: 1": {input: 1, want: Rest},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Categorizer(tc.input)
			if got != tc.want {
				t.Errorf("Categorizer() got = %v, want %v", got, tc.want)
			}
		})
	}
}
