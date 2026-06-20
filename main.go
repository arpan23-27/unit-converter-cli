package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type conversionFunc func(string, float64) (string, error)

func main() {
	if len(os.Args) < 2 || isHelpArg(os.Args[1]) {
		printUsage()
		return
	}

	if len(os.Args) < 4 {
		fail("missing arguments")
	}

	value, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fail(fmt.Sprintf("value must be a number: %s", os.Args[3]))
	}

	switch strings.ToLower(os.Args[1]) {
	case "temp":
		runConversion(convertTemp, os.Args[2], value)
	case "len":
		runConversion(convertLength, os.Args[2], value)
	case "weight":
		runConversion(convertWeight, os.Args[2], value)
	default:
		fail(fmt.Sprintf("unknown category: %s", os.Args[1]))
	}
}

func runConversion(fn conversionFunc, mode string, value float64) {
	out, err := fn(mode, value)
	if err != nil {
		fail(err.Error())
	}
	fmt.Println(out)
}

func isHelpArg(arg string) bool {
	switch strings.ToLower(arg) {
	case "-h", "--help", "help":
		return true
	default:
		return false
	}
}

func fail(message string) {
	fmt.Fprintln(os.Stderr, "error:", message)
	printUsage()
	os.Exit(1)
}

func printUsage() {
	fmt.Println("unit-converter-cli")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  unit-converter <temp|len|weight> <conversion> <value>")
	fmt.Println("  unit-converter --help")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  unit-converter temp c2f 100")
	fmt.Println("  unit-converter temp f2c 32")
	fmt.Println("  unit-converter len m2ft 10")
	fmt.Println("  unit-converter weight kg2lb 5")
}
