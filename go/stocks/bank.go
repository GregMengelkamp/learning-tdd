package stocks

import "errors"

type Bank struct {
	exchangeRates map[string]float64
}

// extends the exchangeRates map of a bank
func (b Bank) AddExchangeRate(currencyFrom string, currencyTo string, rate float64) {
	key := currencyFrom + "->" + currencyTo
	b.exchangeRates[key] = rate
}

// this conversion returns Moneypointer (default go style) and an error not a float64 to indicate an amount and a mere boolean 
func (b Bank) Convert(money Money, currencyTo string) (convertedMoney *Money, err error) {
	var result Money
	if money.currency == currencyTo {
		result = NewMoney(money.amount, money.currency)
		// return of a pointer if no errors occured
		return &result, nil
	}
	key := money.currency + "->" + currencyTo
	rate, ok := b.exchangeRates[key]
	if ok {
		result = NewMoney(money.amount*rate, currencyTo)
		// return of a pointer if no errors occured
		return &result, nil
	}
		return nil, errors.New(key)
	}

// making sure we value dependency injection; hereby creation of a bankstruct is separated from it's use
func NewBank() Bank {
	return Bank{exchangeRates: make(map[string]float64)}
}