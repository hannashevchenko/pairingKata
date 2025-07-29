package categorizer

import (
	"fmt"
	"testing"
)

func TestCategorizer(t *testing.T) {
	tests := []struct {
		input int
		want  Category
	}{
		{input: 1, want: Rest},
		{input: 3, want: Fizz},
		{input: 6, want: Fizz},
		{input: 9, want: Fizz},
		{input: 5, want: Buzz},
		{input: 10, want: Buzz},
		{input: 15, want: FizzBuzz},
	}
	for _, tc := range tests {
		name := fmt.Sprintf("%d", tc.input)
		t.Run(name, func(t *testing.T) {
			got := Categorizer(tc.input)
			if got != tc.want {
				t.Errorf("Categorizer() got = %v, want %v", got, tc.want)
			}
		})
	}
}
