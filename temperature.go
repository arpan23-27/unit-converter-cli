package main

import "fmt"

func cToF(c float64) float64 { return c*9/5 + 32 }
func fToC(f float64) float64 { return (f - 32) * 5 / 9 }

func convertTemp(mode string, v float64) (string, error) {
	switch mode {
	case "c2f":
		return fmt.Sprintf("%.1f C = %.1f F", v, cToF(v)), nil
	case "f2c":
		return fmt.Sprintf("%.1f F = %.1f C", v, fToC(v)), nil
	default:
		return "", fmt.Errorf("unknown temp conversion: %s", mode)
	}
}