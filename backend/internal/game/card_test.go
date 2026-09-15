package game

import "testing"

func TestIsValidCardValue(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{name: "minimum value", value: MinCardValue, want: true},
		{name: "maximum value", value: MaxCardValue, want: true},
		{name: "zero", value: 0, want: true},
		{name: "typical value", value: 7, want: true},
		{name: "below minimum", value: MinCardValue - 1, want: false},
		{name: "above maximum", value: MaxCardValue + 1, want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isValidCardValue(test.value); got != test.want {
				t.Fatalf("isValidCardValue(%d) = %t, want %t", test.value, got, test.want)
			}
		})
	}
}
