package main

import (
	"testing"
)

// test method, must start with Test, must have *testing.T argument
func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{amount: 5, currency: "USD"}
	tenner := fiver.Times(2)
	// actual value vs expected value - definite value!
	if tenner.amount != 10 {
		t.Errorf("Expected 10, got: [%d]", tenner.amount)
	}
	if tenner.currency != "USD" {
		t.Errorf("Expected USD, got: [%s]", tenner.currency)
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

// introducing the new money struct with the necessary properties
type Money struct {
	amount int
	currency string
	}

// creating a corresponding function to multiply money
func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * multiplier, currency: m.currency}
}