const Money = require('./money');

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

module.exports = Portfolio;
