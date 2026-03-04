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

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got != want {

			}
		})
	}
}

