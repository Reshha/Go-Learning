//Package weather provides tool for weather forecast on a specific location.
package weather


var (
    //CurrentCondition is used to define the current weather condition.
	CurrentCondition string
    
    //CurrentLocation is used to define the current location.
	CurrentLocation  string
)

//Forecast returns a string showing the forecast on a city via the city parameter which overrides the value of CurrentLocation and condition via the condition parameter which overrides the value of CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
