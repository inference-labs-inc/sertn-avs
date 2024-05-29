package core

import (
	"fmt"
	"math/big"
	"testing"
)

func TestRandomInputs(t *testing.T) {
	inputs := RandomInputs()
	for _, input := range inputs {
		if input.Cmp(big.NewInt(10)) == 1 {
			t.Errorf("Input greater than 10 from random inputs %s", input.String())
		}
	}
}

func TestFormatInputsForChain(t *testing.T) {
	formattedInputs := FormatInputsForChain([5]float64{0.1, 0.2, 0.3, 0.5, 0.3})
	goodFormattedInputs := [5]int64{0, 1, 1, 2, 1}
	for index, input := range goodFormattedInputs {
		if formattedInputs[index] != big.NewInt(input) {
			t.Errorf("Expected %d to be %s in formatted inputs", input, formattedInputs[index].String())
		}
	}
}

func TestFormatFloatInputsToString(t *testing.T) {
	formattedFloats := FormatFloatInputsToString([5]float64{0.1, 0.2, 0.3, 0.5, 0.3})
	goodFormattedFloats := "0.1 0.2 0.3 0.5 0.3"
	if formattedFloats != goodFormattedFloats {
		t.Errorf("Expected %s to be %s for formatting floats to a string", formattedFloats, goodFormattedFloats)
	}
}

func TestFormatBigIntInputsToString(t *testing.T) {
	formattedInputs := FormatBigIntInputsToString([5]*big.Int{big.NewInt(2), big.NewInt(1), big.NewInt(3), big.NewInt(4), big.NewInt(10)})
	fmt.Println(formattedInputs)
}
