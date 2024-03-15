package lab2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrefixToPostfix1(t *testing.T) {
	res, err := PrefixToPostfix("+ 6 8") 
	if assert.Nil(t, err) {
		assert.Equal(t, "6 8 +", res)
	}
}

func TestPrefixToPostfix2(t *testing.T) {
	res, err := PrefixToPostfix("* - 3 4 2")
	if assert.Nil(t, err) {
		assert.Equal(t, "3 4 - 2 *", res)
	}
}

func TestPrefixToPostfix3(t *testing.T) {
	res, err := PrefixToPostfix("/ * + 1 4 + - 1 4 6 - 2 3")
	if assert.Nil(t, err) {
		assert.Equal(t, "1 4 + 1 4 - 6 + * 2 3 - /", res)
	}
}

func TestPrefixToPostfix4(t *testing.T) {
	res, err := PrefixToPostfix("/ * + - - 2 3 2 7 + + + 2 3 7 3 7")
	if assert.Nil(t, err) {
		assert.Equal(t, "2 3 - 2 - 7 + 2 3 + 7 + 3 + * 7 /", res)
	}
}

func TestPrefixToPostfix5(t *testing.T) {
	_, err := PrefixToPostfix("")

	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("невірний префіксний вираз"), err)
	}
}
func TestPrefixToPostfix6(t *testing.T) {
	_, err := PrefixToPostfix("Плюс три  вісім")

	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("невірний префіксний вираз"), err)
	}
}
func TestPrefixToPostfix7(t *testing.T) {
	_, err := PrefixToPostfix("-....- ..... ...--")

	assert.NotNil(t, err)
	if assert.Error(t, err) {
		assert.Equal(t, fmt.Errorf("невірний префіксний вираз"), err)
	}
}

// Example demonstrates how to use PrefixToPostfix function.
func ExamplePrefixToPostfix() {
	res, _ := PrefixToPostfix("- 5 8")
	fmt.Println(res)

	// Output:
	// 5 8 -
}
