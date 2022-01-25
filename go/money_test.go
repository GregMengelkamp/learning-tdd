package main

import (
	"testing"
	s "tdd/stocks"
)

func assertEqual(t *testing.T, expected s.Money, actual s.Money) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

// test method, must start with Test, must have *testing.T argument
func TestMultiplication(t *testing.T) {
	tenEuros := s.NewMoney(10, "EUR")
	actualResult := tenEuros.Times(2)
	expectedResult := s.NewMoney(20, "EUR")
	assertEqual(t, expectedResult, actualResult)
}

func TestDivision(t *testing.T) {
	originalMoney := s.NewMoney(4002, "KRW")
	expectedResult := originalMoney.Divide(4)
	actualResult := s.NewMoney(1000.5, "KRW")
	assertEqual(t, expectedResult, actualResult)
}

// doesn't it look like pascal a lot?! ":="
func TestAddition(t *testing.T) {
	// empty portfolio; emphasizing expected types!
	var portfolio s.Portfolio
	var portfolioInDollars s.Money

	fiveDollars := s.NewMoney(5,"USD")
	tenDollars := s.NewMoney(10, "USD")

	// money entity to compare to
	fifteenDollars := s.NewMoney(15, "USD")

	// adding money to portfolio
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)

	// portfolio evaluate should return a money struct of specified currency
	portfolioInDollars = portfolio.Evaluate("USD")
	assertEqual(t, fifteenDollars, portfolioInDollars)
}



