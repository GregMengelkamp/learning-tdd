import unittest

from money import Money
from portfolio import Portfolio

# creating testclass inherting from unittest.TestCase - soweit so bekannt aus div. Videos
class TestMoney(unittest.TestCase):
    # testcase for the new Money class
    def testMultiplication(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = Money(20, "EUR")
        twentyEuros = tenEuros.times(2)
        self.assertEqual(twentyEuros, tenEuros.times(2))

    def testDivision(self):
        originalMoney = Money(4002, "KRW")
        actualMoneyAfterDivision = originalMoney.divide(4)
        expectedMoneyAfterDivision = Money(1000.5, "KRW")
        self.assertEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision)

    def testAddition(self):
        fiveDollars = Money(5, "USD")
        tenDollars = Money(10, "USD")
        fifteenDollars = Money(15, "USD")
        portfolio = Portfolio()
        portfolio.add(fiveDollars, tenDollars)
        self.assertEqual(fifteenDollars, portfolio.evaluate("USD"))

    # assuming a unirectional exchangerate of 1.2USD/EUR
    def testAdditionOfDollarsAndEuros(self):
        fiveDollars = Money(5, "USD")
        tenEuros = Money(10, "EUR")
        portfolio = Portfolio()
        portfolio.add(fiveDollars, tenEuros)
        expectedValue = Money(17, "USD")
        actualValue = portfolio.evaluate("USD")
        self.assertEqual(
            expectedValue, actualValue, "%s != %s" % (expectedValue, actualValue)
        )


# run all tests if ran as main
if __name__ == "__main__":
    unittest.main()
