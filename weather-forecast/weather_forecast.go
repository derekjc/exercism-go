// Package weather provides Forecasting capabilities.
package weather

// CurrentCondition ... 
var CurrentCondition string
// CurrentLocation ... else.
var CurrentLocation string

// Forecast returns a string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
