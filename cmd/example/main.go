package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	lab2 "github.com/vladimirkoff/arch-lab2"
)

var (
	inputExpression = flag.String("e", "", "Expression to compute")
	inputFile       = flag.String("f", "", "Input file")
	outputFile      = flag.String("o", "", "Output file")
)

func main() {
	flag.Parse()

	var input io.Reader

	if *inputExpression != "" && *inputFile != "" {
		fmt.Fprintln(os.Stderr, "забагато прапорів")
	}

	if *inputExpression != "" {
		input = strings.NewReader(*inputExpression)
	} else if *inputFile != "" {
		file, err := os.Open(*inputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "поилка з файлом результату: "+err.Error())
		}
		defer file.Close()
		input = file
	} else {
		fmt.Fprintln(os.Stderr, "помилка з отриманим параметром")
		return
	}

	var res io.Writer
	if *outputFile != "" {
		file, err := os.Create(*outputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "помилка результату: "+err.Error())
		}
		defer file.Close()
		res = file
	} else {
		res = os.Stdout
	}

	handler := lab2.ComputeHandler{
		Input:  input,
		Output: res,
	}

	err := handler.Compute()
	if err != nil {
		fmt.Fprintln(os.Stderr, "помилка конвертації: "+err.Error())
	}
}