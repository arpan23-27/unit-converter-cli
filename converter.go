package main

import (
	"fmt"
	"strings"
)

type conversionFunc func(string, float64) (string, error)

func convertTemp(mode string, value float64) (string, error) {
	switch strings.ToLower(mode) {
	case "c2f":
		return fmt.Sprintf("%.1f C = %.1f F", value, celsiusToFahrenheit(value)), nil
	case "f2c":
		return fmt.Sprintf("%.1f F = %.1f C", value, fahrenheitToCelsius(value)), nil
	default:
		return "", fmt.Errorf("unknown temperature conversion: %s", mode)
	}
}

func convertLength(mode string, value float64) (string, error) {
	switch strings.ToLower(mode) {
	case "m2ft":
		return fmt.Sprintf("%.1f m = %.2f ft", value, metersToFeet(value)), nil
	case "ft2m":
		return fmt.Sprintf("%.1f ft = %.2f m", value, feetToMeters(value)), nil
	default:
		return "", fmt.Errorf("unknown length conversion: %s", mode)
	}
}

func convertWeight(mode string, value float64) (string, error) {
	switch strings.ToLower(mode) {
	case "kg2lb":
		return fmt.Sprintf("%.1f kg = %.2f lb", value, kilogramsToPounds(value)), nil
	case "lb2kg":
		return fmt.Sprintf("%.1f lb = %.2f kg", value, poundsToKilograms(value)), nil
	default:
		return "", fmt.Errorf("unknown weight conversion: %s", mode)
	}
}

func celsiusToFahrenheit(value float64) float64 {
	return value*9/5 + 32
}

func fahrenheitToCelsius(value float64) float64 {
	return (value - 32) * 5 / 9
}

func metersToFeet(value float64) float64 {
	return value * 3.28084
}

func feetToMeters(value float64) float64 {
	return value / 3.28084
}

func kilogramsToPounds(value float64) float64 {
	return value * 2.20462
}

func poundsToKilograms(value float64) float64 {
	return value / 2.20462
}
