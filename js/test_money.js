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
	divide(divisor) {
		return new Money(this.amount / divisor, this.currency);
	}
}

// creating fiver object
let fiver = new Money(5, 'USD');
let tenner = new Money(10, 'USD');
// actual value vs expected value - strictEqual assert statement
assert.deepStrictEqual(fiver.times(2), tenner);

// new testcase for money object with amount and currency
let tenEuros = new Money(10, 'EUR');
let twentyEuros = new Money(20, 'EUR');
assert.deepStrictEqual(tenEuros.times(2), twentyEuros);

// adds testcase for divison
let originalMoney = new Money(4002, 'KRW');
let actualMoneyAfterDivision = originalMoney.divide(4);
let expectedMoneyAfterDivision = new Money(1000.5, 'KRW');
assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);
