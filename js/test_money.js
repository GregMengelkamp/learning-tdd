// loading in the assert package to be able to test
const assert = require('assert');

// implementing a Dollar class, initialize BEFORE using it
class Dollar {
	constructor(amount) {
		this.amount = amount;
	}
	// times function with multiplier parameter -> w/o explicit return type gives undefined
	times(multiplier) {
		// from hardcoded ten to "destructured" 10; revealing the abstractions -> multiplication, removed duplication and coupling
		return new Dollar(this.amount * multiplier);
	}
}

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
let fiver = new Dollar(5);
// creating tenner object out of fiver object with times-method
let tenner = fiver.times(2);

// actual value vs expected value - strictEqual assert statement
assert.strictEqual(tenner.amount, 10);

// new testcase for money object with amount and currency
let tenEuros = new Money(10, 'EUR');
let twentyEuros = tenEuros.times(2);
assert.strictEqual(twentyEuros.amount, 20);
assert.strictEqual(twentyEuros.currency, 'EUR');
