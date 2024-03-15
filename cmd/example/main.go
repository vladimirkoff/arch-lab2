package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	lab2 "github.com/vladimirkoff/arch-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	input, err := getInputReader()
	if err != nil {
		fmt.Fprintln(os.Stderr, "input failed:", err)
		return
	}

	output, err := getOutputWriter()
	if err != nil {
		fmt.Fprintln(os.Stderr, "output failed:", err)
		return
	}

	handler := lab2.ComputeHandler{
		Input:  input,
		Output: output,
	}
	if err := handler.Compute(); err != nil {
		fmt.Fprintln(os.Stderr, "computation failed:", err)
	}
}

func getInputReader() (io.Reader, error) {
	if *inputExpression != "" && *inputFile != "" {
		return nil, fmt.Errorf("exceeded amount of input flags")
	}

	if *inputExpression != "" {
		return strings.NewReader(*inputExpression), nil
	}

	if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			return nil, fmt.Errorf("input file error: %s", err)
		}
		defer file.Close()
		return file, nil
	}
	return nil, fmt.Errorf("no input")
}

func getOutputWriter() (io.Writer, error) {
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			return nil, fmt.Errorf("output file error: %s", err)
		}
		defer file.Close()
		return file, nil
	}
	return os.Stdout, nil
}
