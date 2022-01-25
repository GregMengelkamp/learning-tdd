package main

import (
	"testing"
)

func assertEqual(t *testing.T, expected Money, actual Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

// test method, must start with Test, must have *testing.T argument
func TestMultiplicationInDollars(t *testing.T) {
	fiver := Money{amount: 5, currency: "USD"}
	actualResult := fiver.Times(2)
	expectedResult := Money{amount: 10, currency: "USD"}
	// actual value vs expected value - definite value!
	assertEqual(t, expectedResult, actualResult)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualResult := tenEuros.Times(2)
	expectedResult := Money{amount: 20, currency: "EUR"}
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := Money{amount: 4002, currency: "KRW"}
	expectedResult := originalMoney.Divide(4)
	actualResult := Money{amount: 1000.5, currency: "KRW"}
	assertEqual(t, expectedResult, actualResult)
}

// doesn't it look like pascal a lot?! ":="
func TestAddition(t *testing.T) {
	// empty portfolio; emphasizing expected types!
	var portfolio Portfolio
	var portfolioInDollars Money

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}

	// money entity to compare to
	fifteenDollars := Money{amount: 15, currency: "USD"}

	// adding money to portfolio
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)

	// portfolio evaluate should return a money struct of specified currency
	portfolioInDollars = portfolio.Evaluate("USD")
	assertEqual(t, fifteenDollars, portfolioInDollars)
}

// introducing the new money struct with the necessary properties
type Money struct {
	amount   float64
	currency string
}

// creating a corresponding function to multiply money
func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}

// an array of money entities
type Portfolio []Money
// returns the portfolio itself atm
// append ( to what, what to append )
func (p Portfolio) Add(money Money) Portfolio {
	p = append(p, money)
	return p
}

// returns the hardcoded correct amount (least amount of code posssible to pass...)
// returns actual amounts; iteration a mix of enumerate python & strong pascal resemblance
func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, m := range p {
		total = total + m.amount
	}
	return Money{amount: total, currency: currency}
}