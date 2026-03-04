package main

import "testing"

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name string
		input float64
		want float64
	}{
		{name: "normal case", input: 13.0, want: (13.0 - 32.0) * 5/9},
		{name: "negative", input: -65.6, want: (-65.6 - 32.0) * 5/9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FahrenheitToCelsius(test.input)
			if got != want {
				t.Errorf("input: %v, got: %v, want: %v", test.input, got, test.want)
			}
		})
	}
}

