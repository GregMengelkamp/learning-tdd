import unittest


class Dollar:
    def __init__(self, amount):
        self.amount = amount

    def times(self, multiplier):
        return Dollar(self.amount * multiplier)


class Money:
    def __init__(self, amount, currency):
        self.amount = amount
        self.currency = currency

    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)


# creating testclass inherting from unittest.TestCase - soweit so bekannt aus div. Videos
class TestMoney(unittest.TestCase):
    # first test(+method) - create fiver & tenner
    def testMultiplication(self):
        fiver = Dollar(5)
        tenner = fiver.times(2)

        # actual value vs expected value with assertEqual expression from unittest
        self.assertEqual(10, tenner.amount)

    # testcase for the new Money class
    def testMultiplicationInEuros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = tenEuros.times(2)
        self.assertEqual(20, twentyEuros.amount)
        self.assertEqual("EUR", twentyEuros.currency)


# run all tests if ran as main
if __name__ == "__main__":
    unittest.main()
