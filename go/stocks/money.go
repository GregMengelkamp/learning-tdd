package stocks

// encapsulated in stocks package -> immutable from outside
type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}

// amount and currency fields of Money struct are accessible to the public by this function
func NewMoney(amount float64, currency string) Money {
	return Money{amount, currency}
}