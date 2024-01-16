package main

import "fmt"

func main() {
	var atmTemperature float64
	var atmPressure float64
	var atmHumidity float64

	atmTemperature = 9.6
	atmPressure = 775.9
	atmHumidity = 68.9

	fmt.Println("Atmospheric Temperature:", atmTemperature, "°C")
	fmt.Println("Atmospheric Pressure:", atmPressure, "%")
	fmt.Println("Atmospheric Humidity:", atmHumidity, "hPa")
}
