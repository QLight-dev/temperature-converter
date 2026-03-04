package main

import "testing"

func TestFahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{name: "normal case", input: 13.0, want: (13.0 - 32.0) * 5 / 9},
		{name: "negative", input: -65.6, want: (-65.6 - 32.0) * 5 / 9},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FahrenheitToCelsius(test.input)

			// floating-point arithmetic won't be exact, so we check if its close enough to the expected value
			delta := 0.0001
			if got < test.want-delta || got > test.want+delta {
				t.Errorf("input: %v, got: %v, want: %v", test.input, got, test.want)
			}
		})
	}
}

func TestFahreheitToKelvin(t *testing.T) {
	tests := []struct {
		name  string
		input float64
		want  float64
	}{
		{name: "normal case", input: 98.5, want: (98.5-32.0)*5/9 + 273.15},
		{name: "negative", input: -75.6, want: (-75.6-32.0)*5/9 + 273.15},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := FahreheitToKelvin(test.input)

			// floating-point arithmetic won't be exact, so we check if its close enough to the expected value
			delta := 0.0001
			if got < test.want-delta || got > test.want+delta {
				t.Errorf("input: %v, got: %v, want: %v", test.input, got, test.want)
			}
		})
	}
}
