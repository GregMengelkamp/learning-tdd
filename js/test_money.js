// loading in the assert package to be able to test
const assert = require('assert');

// creating fiver object
let fiver = new Dollar(5);
// creating tenner object out of fiver object with times-method
let tenner = fiver.times(2);

// actual value vs expected value - strictEqual assert statement
assert.strictEqual(tenner.amount, 10);