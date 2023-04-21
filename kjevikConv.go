package kjevikConv

import (
	"log"
	"strconv"
	"strings"
)

func celsiusToFahrenheit(temp float64) float64 {

	return (temp * 9 / 5) + 32

}

func konverterSetning(input string) string {

	parts := strings.Split(input, ";")

	if len(parts) != 4 {
		log.Println("Please input valid data")
	}

	tempvalue, err := strconv.ParseFloat(parts[3], 64)
	if err != nil || parts[3] == "" {
		log.Println("Error parsing temperature:", err)
	}

	fahrenheit := celsiusToFahrenheit(tempvalue)
	fahrenheitString := strconv.FormatFloat(fahrenheit, 'f', 1, 64)

	parts[3] = fahrenheitString

	konvertertSetning := strings.Join(parts, ";")

	return konvertertSetning

}
