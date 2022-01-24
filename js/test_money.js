// loading in the assert package to be able to test
const assert = require('assert');

// implements Money class, initializes BEFORE using it
class Money {
	constructor(amount, currency) {
		this.amount = amount;
		this.currency = currency;
	}
	times(multiplier) {
		return new Money(this.amount * multiplier, this.currency);
	}
}

// creating fiver object
let fiver = new Money(5, 'USD');
let tenner = fiver.times(2);

// actual value vs expected value - strictEqual assert statement
assert.strictEqual(tenner.amount, 10);
assert.strictEqual(tenner.currency, 'USD');

// new testcase for money object with amount and currency
let tenEuros = new Money(10, 'EUR');
let twentyEuros = tenEuros.times(2);
assert.strictEqual(twentyEuros.amount, 20);
assert.strictEqual(twentyEuros.currency, 'EUR');
