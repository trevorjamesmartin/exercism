// Package weather can forecast the current weather condition of various cities in Goblinocus.
package weather

// CurrentCondition stores a description of the weather condition at CurrentLocation.
var CurrentCondition string

// CurrentLocation stores the name of a city or province in the country of Goblinocus.
var CurrentLocation string

// Forecast returns real-time weather data for the citizens of Goblinocus.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
