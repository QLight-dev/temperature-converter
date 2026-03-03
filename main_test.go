package main

import "testing"

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name string
		got float64
		want float64
	}{
		{name: "normal case", got: FahrenheitToCelsius(13.0), want: (13.0 - 32.0) * 5/9},
		{name: "negative", got: FahrenheitToCelsius(-65.0), want: (-65.0 - 32.0) * 5/9},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if got != want {

			}
		})
	}
}

