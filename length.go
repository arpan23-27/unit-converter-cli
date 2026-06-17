package main

import "fmt"

func convertLength(mode string, v float64) (string, error) {
	switch mode {
	case "m2ft":
		return fmt.Sprintf("%.1f m = %.2f ft", v, v*3.28084), nil
	case "ft2m":
		return fmt.Sprintf("%.1f ft = %.2f m", v, v/3.28084), nil
	default:
		return "", fmt.Errorf("unknown length conversion: %s", mode)
	}
}