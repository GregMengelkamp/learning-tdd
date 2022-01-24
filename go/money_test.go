package main

import (
	"testing"
)

// test method, must start with Test, must have *testing.T argument
func TestMultiplication(t *testing.T) {
	// fünfer ist ein Dollar mal fünf
	fiver := Dollar{
		amount: 5,
	}
	// Zehner ist Fünfer (struct/object) mal zwei
	tenner := fiver.Times(2)
	// actual value vs expected value - definite value!
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%d]", tenner.amount)
	}
}

// introducing a Dollar struct
type Dollar struct {
	// adding amount field
	amount int
	}

/* a functions of a Dollar named Times which takes in an int multiplier and return a Dollar struct 
whose amount is equal to multiplier * amount of the original dollar struct 
Kent Beck - "If dependency is the problem, duplication is the symptom."*/
func (d Dollar) Times(multiplier int) Dollar {
	return Dollar{amount: d.amount * multiplier}
	}
