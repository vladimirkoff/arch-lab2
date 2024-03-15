package lab2

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompute(t *testing.T) {
	type examples struct {
		name           string
		input          string
		expectedOutput string
		expectedError  error
	}

	var examplesTable = []examples{
		{
			name:           "Example 1",
			input:          "- - 7 5 3",
			expectedOutput: "7 5 - 3 -",
			expectedError:  nil,
		},
		{
			name:           "Example 2",
			input:          " ",
			expectedOutput: "",
			expectedError:  errors.New("Невірний префіксний вираз"),
		},
		{
			name:           "Example 3",
			input:          "...--- .-.-.- ---...",
			expectedOutput: "",
			expectedError:  errors.New("Невірний префіксний вираз"),
		},
	}

	for _, test := range examplesTable {
		input := strings.NewReader(test.input)
		output := new(bytes.Buffer)

		handler := ComputeHandler{
			Input:  input,
			Output: output,
		}

		err := handler.Compute()

		outputToString := output.String()

		assert.Equal(t, test.expectedOutput, outputToString, test.name)
		assert.Equal(t, test.expectedError, err, test.name)
	}
}
