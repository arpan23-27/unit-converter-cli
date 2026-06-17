package main

import "fmt"

func convertWeight(mode string, v float64) (string, error) {
	switch mode {
	case "kg2lb":
		return fmt.Sprintf("%.1f kg = %.2f lb", v, v*2.20462), nil
	case "lb2kg":
		return fmt.Sprintf("%.1f lb = %.2f kg", v, v/2.20462), nil
	default:
		return "", fmt.Errorf("unknown weight conversion: %s", mode)
	}
}