package main

import "testing"

func TestTemperatureHelpers(t *testing.T) {
	if got := cToF(100); got != 212 {
		t.Fatalf("cToF(100) = %v, want 212", got)
	}
	if got := fToC(32); got != 0 {
		t.Fatalf("fToC(32) = %v, want 0", got)
	}
}

func TestConvertTemp(t *testing.T) {
	got, err := convertTemp("c2f", 100)
	if err != nil {
		t.Fatalf("convertTemp returned error: %v", err)
	}
	want := "100.0 C = 212.0 F"
	if got != want {
		t.Fatalf("convertTemp(c2f) = %q, want %q", got, want)
	}
}

func TestConvertLength(t *testing.T) {
	got, err := convertLength("m2ft", 10)
	if err != nil {
		t.Fatalf("convertLength returned error: %v", err)
	}
	want := "10.0 m = 32.81 ft"
	if got != want {
		t.Fatalf("convertLength(m2ft) = %q, want %q", got, want)
	}
}

func TestConvertWeight(t *testing.T) {
	got, err := convertWeight("kg2lb", 5)
	if err != nil {
		t.Fatalf("convertWeight returned error: %v", err)
	}
	want := "5.0 kg = 11.02 lb"
	if got != want {
		t.Fatalf("convertWeight(kg2lb) = %q, want %q", got, want)
	}
}
