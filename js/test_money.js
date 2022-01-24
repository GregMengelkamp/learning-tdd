// loading in the assert package to be able to test
const assert = require('assert');

// implementing a Dollar class, initialize BEFORE using it
class Dollar {
    constructor(amount) {
        this.amount = amount;
    }
    // times function with multiplier parameter -> w/o explicit return type gives undefined
    times(multiplier) {
        return new Dollar(10)
    }
}

// creating fiver object
let fiver = new Dollar(5);
// creating tenner object out of fiver object with times-method
let tenner = fiver.times(2);

// actual value vs expected value - strictEqual assert statement
assert.strictEqual(tenner.amount, 10);
