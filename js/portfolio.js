const Money = require('./money');

class Portfolio {
	constructor() {
		this.moneys = [];
	}
	// rest parameter syntax allows a function to accept an indefinite number of arguments as an array
	add(...moneys) {
		this.moneys = this.moneys.concat(moneys);
	}

	evaluate(bank, currency) {
		let failures = [];
		let total = this.moneys.reduce((sum, money) => {
			try {
				let convertedMoney = bank.convert(money, currency);
				return sum + convertedMoney.amount;
			} catch (error) {
				failures.push(error.message);
				return sum;
			}
		}, 0);
		// no elements in an array <-> !array.length
		if (!failures.length) {
			return new Money(total, currency);
		}
		throw new Error('Missing exchange rate(s):[' + failures.join() + ']');
	}
}

module.exports = Portfolio;
