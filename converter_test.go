package main

import "testing"

func TestTemperature(t *testing.T) {
	if celsiusToFahrenheit(100) != 212 {
		t.Fatal("100C should be 212F")
	}
	if fahrenheitToCelsius(32) != 0 {
		t.Fatal("32F should be 0C")
	}
}

func TestLength(t *testing.T) {
	if metersToFeet(10) < 32.8 || metersToFeet(10) > 32.81 {
		t.Fatal("unexpected meter to feet conversion")
	}
}

func TestWeight(t *testing.T) {
	if kilogramsToPounds(5) < 11 || kilogramsToPounds(5) > 11.1 {
		t.Fatal("unexpected kilogram to pound conversion")
	}
}
