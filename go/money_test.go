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
