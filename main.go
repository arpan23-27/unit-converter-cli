package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
	switch os.Args[1] {
	case "temp":
		runConvert(convertTemp)
	case "len":
		runConvert(convertLength)
	default:
		fmt.Println("unknown command:", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

func runConvert(fn func(string, float64) (string, error)) {
	if len(os.Args) < 4 {
		printUsage()
		os.Exit(1)
	}
	value, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("value must be a number:", os.Args[3])
		os.Exit(1)
	}
	out, err := fn(os.Args[2], value)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(out)
}

func printUsage() {
	fmt.Println("Usage: unit-converter <temp|len|weight> <conversion> <value>")
}