package core

import (
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
