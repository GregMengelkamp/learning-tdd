const assert = require('assert');
const Money = require('./money');
const Portfolio = require('./portfolio');

class MoneyTest {
	testMultiplication() {
		let tenEuros = new Money(10, 'EUR');
		let twentyEuros = new Money(20, 'EUR');
		assert.deepStrictEqual(tenEuros.times(2), twentyEuros);
	}

	testDivision() {
		let originalMoney = new Money(4002, 'KRW');
		let actualMoneyAfterDivision = originalMoney.divide(4);
		let expectedMoneyAfterDivision = new Money(1000.5, 'KRW');
		assert.deepStrictEqual(
			actualMoneyAfterDivision,
			expectedMoneyAfterDivision,
		);
	}

	testAddition() {
		let fiveDollars = new Money(5, 'USD');
		let tenDollars = new Money(10, 'USD');
		let fifteenDollars = new Money(15, 'USD');
		let portfolio = new Portfolio();
		portfolio.add(fiveDollars, tenDollars);
		assert.deepStrictEqual(portfolio.evaluate('USD'), fifteenDollars);
	}

	// assuming a unirectional exchangerate of 1.2USD/EUR
	testAdditionOfDollarsAndEuros() {
		let fiveDollars = new Money(5, 'USD');
		let tenEuros = new Money(10, 'EUR');
		let portfolio = new Portfolio();
		portfolio.add(fiveDollars, tenEuros);
		let expectedValue = new Money(17, 'USD');
		assert.deepStrictEqual(portfolio.evaluate('USD'), expectedValue);
	}

	// assuming a unirectional exchangerate of 1100KRW/USD
	testAdditionOfDollarsAndWons() {
		let oneDollar = new Money(1, 'USD');
		let elevenHundredWon = new Money(1100, 'KRW');
		let portfolio = new Portfolio();
		portfolio.add(oneDollar, elevenHundredWon);
		let expectedValue = new Money(2200, 'KRW');
		assert.deepStrictEqual(portfolio.evaluate('KRW'), expectedValue);
	}

	testAdditionWithMultipleMissingExchangeRates() {
		let oneDollar = new Money(1, 'USD');
		let oneEuro = new Money(1, 'EUR');
		let oneWon = new Money(1, 'KRW');
		let portfolio = new Portfolio();
		portfolio.add(oneDollar, oneEuro, oneWon);
		let expectedError = new Error(
			'Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid]',
		);
		// assert.throws takes reference to an anonymous function that calls the evaluate
		assert.throws(function () {
			portfolio.evaluate('Kalganid');
		}, expectedError);
	}

	// get all (own )properties of a MoneyTest prototype
	getAllTestMethods() {
		let moneyPrototype = MoneyTest.prototype;
		let allProps = Object.getOwnPropertyNames(moneyPrototype);
		// filter all properties for being a function and beginning with test
		let testMethods = allProps.filter((p) => {
			return typeof moneyPrototype[p] === 'function' && p.startsWith('test');
		});
		return testMethods;
	}

	// get all test*method names, create method in relation to THIS class, run the methods with respect to this class w/o arguments
	runAllTests() {
		let testMethods = this.getAllTestMethods();
		testMethods.forEach((m) => {
			let method = Reflect.get(this, m);
			console.log('Running: %s()', m);
			try {
				Reflect.apply(method, this, []);
			} catch (e) {
				if (e instanceof assert.AssertionError) {
					console.log(e);
				} else {
					throw e;
				}
			}
		});
	}
}

new MoneyTest().runAllTests();
