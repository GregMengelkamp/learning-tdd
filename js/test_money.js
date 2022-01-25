// loading in the assert package to be able to test
const assert = require('assert');
const { time } = require('console');

// implements Money class, initializes BEFORE using it
class Money {
	constructor(amount, currency) {
		this.amount = amount;
		this.currency = currency;
	}
	times(multiplier) {
		return new Money(this.amount * multiplier, this.currency);
	}
	divide(divisor) {
		return new Money(this.amount / divisor, this.currency);
	}
}

class Portfolio {
	constructor() {
		this.moneys = [];
	}
	// rest parameter syntax allows a function to accept an indefinite number of arguments as an array
	add(...moneys) {
		this.moneys = this.moneys.concat(moneys);
	}
	evaluate(currency) {
		// using reduce outside of webdev or fn-examples in JS for the first time, i feel; moneys is missing
		let total = this.moneys.reduce((sum, money) => {
			return sum + money.amount;
		}, 0);
		return new Money(total, currency);
	}
}

// creating fiver object
let fiveDollars = new Money(5, 'USD');
let tenDollars = new Money(10, 'USD');
// actual value vs expected value - strictEqual assert statement
assert.deepStrictEqual(fiveDollars.times(2), tenDollars);

// new testcase for money object with amount and currency
let tenEuros = new Money(10, 'EUR');
let twentyEuros = new Money(20, 'EUR');
assert.deepStrictEqual(tenEuros.times(2), twentyEuros);

// adds testcase for divison
let originalMoney = new Money(4002, 'KRW');
let actualMoneyAfterDivision = originalMoney.divide(4);
let expectedMoneyAfterDivision = new Money(1000.5, 'KRW');
assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);

// portfoliotest; empty portfolio, adding money entities of same crc, deepStrictEqual of portfolio.evaluate and expected money entity
let fifteenDollars = new Money(15, 'USD');
let portfolio = new Portfolio();
portfolio.add(fiveDollars, tenDollars);
assert.deepStrictEqual(portfolio.evaluate('USD'), fifteenDollars);
