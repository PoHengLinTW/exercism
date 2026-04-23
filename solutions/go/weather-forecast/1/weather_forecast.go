// Package weather provides forecast function as well as the current condition and location.
package weather

var (
    // CurrentCondition represents the weather condition for current forecasting.
	CurrentCondition string
    // CurrentLocation represents the city for current forecasting.
	CurrentLocation  string
)

// Forecast the weather condition in a city in a sentence. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
