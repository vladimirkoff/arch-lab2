package lab2

import (
	"bufio"
	"io"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (compHandler *ComputeHandler) Compute() error {
	scanner := bufio.NewScanner(compHandler.Input)
	scanner.Scan()

	inputExpression := scanner.Text()

	res, err := PrefixToPostfix(inputExpression)
	if err != nil {
		return err
	}

	_, err = compHandler.Output.Write([]byte(res))
	if err != nil {
		return err
	}

	return nil
}
