package main

const (
	freezingF = 32.0
	factor    = 5.0 / 9.0
)

// converts Fahrenheit temperatures into Celsius
func FahrenheitToCelsius(temperature float64) float64 {
	return (temperature - freezingF) * factor
}

const (
	kelvinAbsoluteZeroPoint = 273.15
)

// converts Fahrenheit temperatures into Kelvin
func FahreheitToKelvin(temperature float64) float64 {
	return FahrenheitToCelsius(temperature) + kelvinAbsoluteZeroPoint
}
