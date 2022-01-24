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

// this new test introduces a more abstract money class which has a currency field that can be assigned
func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	twentyEuros := tenEuros.Times(2)
	if twentyEuros.amount != 20 {
		t.Errorf("Expected 20, got: [%d]", twentyEuros.amount)
	}
	if twentyEuros.currency != "EUR" {
		t.Errorf("Expected EUR, got: [%s]", twentyEuros.currency)
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

// introducing the new money struct with the necessary properties
type Money struct {
	amount int
	currency string
	}

// creating a corresponding function to multiply money
func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}