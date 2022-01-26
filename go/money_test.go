package main

import (
	s "tdd/stocks"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
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

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")

	// money entity to compare to
	fifteenDollars := s.NewMoney(15, "USD")

	// adding money to portfolio
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)

	// portfolio evaluate should return a money struct of specified currency
	portfolioInDollars, _ = portfolio.Evaluate("USD")
	assertEqual(t, fifteenDollars, portfolioInDollars)
}

// EUR -> USD with an exchange rate of 1.2USD/EUR
func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio
	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := s.NewMoney(17, "USD")
	actualValue, _ := portfolio.Evaluate("USD")
	assertEqual(t, expectedValue, actualValue)
}

// assuming 1100 KRW for every 1 USD
func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio s.Portfolio
	oneDollar := s.NewMoney(1, "USD")
	elevenHundredWon := s.NewMoney(1100, "KRW")
	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)
	expectedValue := s.NewMoney(2200, "KRW")
	actualValue, _ := portfolio.Evaluate("KRW")
	assertEqual(t, expectedValue, actualValue)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	var portfolio s.Portfolio
	oneDollar := s.NewMoney(1, "USD")
	oneEuro := s.NewMoney(1, "EUR")
	oneWon := s.NewMoney(1, "KRW")
	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)
	expectedErrorMessage :=
		"Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"

	// We don’t care about the first return value, so we assign it to the blank identifier.
	_, actualError := portfolio.Evaluate("Kalganid")

	assertEqual(t, expectedErrorMessage, actualError.Error())
	}
