package main

const (
	freezingF = 32
	factor    = 5.0 / 9.0
)

// converts Fahrenheit temperatures into Celsius
func FahrenheitToCelsius(temperature float64) float64 {
	return (temperature - freezingF) * factor
}
