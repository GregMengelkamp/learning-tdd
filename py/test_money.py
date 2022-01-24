import unittest


class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency

    def __eq__(self, other):
        return self.amount == other.amount and self.currency == other.currency

    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)

    def divide(self, divisor):
        return Money(self.amount / divisor, self.currency)


# creating testclass inherting from unittest.TestCase - soweit so bekannt aus div. Videos
class TestMoney(unittest.TestCase):
    # first test(+method) - create fiver & tenner
    def testMultiplicationInDollars(self):
        fiver = Money(5, "USD")
        tenner = Money(10, "USD")

        # actual value vs expected value with assertEqual expression from unittest
        self.assertEqual(tenner, fiver.times(2))

    # testcase for the new Money class
    def testMultiplicationInEuros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = Money(20, "EUR")
        twentyEuros = tenEuros.times(2)
        self.assertEqual(twentyEuros, tenEuros.times(2))

    def testDivision(self):
        originalMoney = Money(4002, "KRW")
        actualMoneyAfterDivision = originalMoney.divide(4)
        expectedMoneyAfterDivision = Money(1000.5, "KRW")
        self.assertEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision)


# run all tests if ran as main
if __name__ == "__main__":
    unittest.main()
