package main

import (
	"reflect"
	s "tdd/stocks"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %+v Got %+v", expected, actual)
	}
}

var bank s.Bank
func init() {
	bank = s.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
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

	fiveDollars := s.NewMoney(5, "USD")
	tenDollars := s.NewMoney(10, "USD")

	// money entity to compare to
	fifteenDollars := s.NewMoney(15, "USD")

	// adding money to portfolio
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)

	// portfolio evaluate should return a money pointer and an error
	portfolioInDollars, err := portfolio.Evaluate(bank, "USD")	
	// there shouldn't be conversion error and the outcome should be 15$	
	assertNil(t, err)
	assertEqual(t, fifteenDollars, *portfolioInDollars)

}

// EUR -> USD with an exchange rate of 1.2USD/EUR
func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio s.Portfolio
	fiveDollars := s.NewMoney(5, "USD")
	tenEuros := s.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	expectedValue := s.NewMoney(17, "USD")

	portfolioInDollars, err := portfolio.Evaluate(bank, "USD")		
	assertNil(t, err)
	assertEqual(t, expectedValue, *portfolioInDollars)
}

// assuming 1100 KRW for every 1 USD
func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio s.Portfolio
	oneDollar := s.NewMoney(1, "USD")
	elevenHundredWon := s.NewMoney(1100, "KRW")
	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)
	expectedValue := s.NewMoney(2200, "KRW")

	portfolioInDollars, err := portfolio.Evaluate(bank, "KRW")		
	assertNil(t, err)
	assertEqual(t, expectedValue, *portfolioInDollars)
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
	value, actualError := portfolio.Evaluate(bank, "Kalganid")
	assertNil(t, value)
	assertEqual(t, expectedErrorMessage, actualError.Error())
}

func TestConversion(t *testing.T) {
	bank := s.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	tenEuros := s.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, s.NewMoney(12, "USD"), *actualConvertedMoney)
}

func assertNil(t *testing.T, actual interface{}) {
	// make sure actual is not a nil interface; && the value of the interface is not NIL, throw error
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Expected to be nil, found: [%+v]", actual)
	}
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	bank := s.NewBank()
	tenEuros := s.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "Kalganid")
	assertNil(t, actualConvertedMoney)
	assertEqual(t, "EUR->Kalganid", err.Error())
}
