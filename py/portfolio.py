import functools
import operator
from money import Money


class Portfolio:
    def __init__(self):
        self.moneys = []

    def add(self, *moneys):
        # should have another look at common methods for iterables..."extend list by appending elements from iterable"
        self.moneys.extend(moneys)

    def evaluate(self, currency):
        # map the moneys array to an array of all amounts, reduce the new array by adding; rather functional python approach
        total = functools.reduce(
            operator.add, map(lambda m: self.__convert(m, currency), self.moneys), 0
        )
        return Money(total, currency)

    def __convert(self, aMoney, aCurrency):
        exchangeRates = {"EUR->USD": 1.2, "USD->KRW": 1100}
        if aMoney.currency == aCurrency:
            return aMoney.amount
        else:
            key = aMoney.currency + "->" + aCurrency
            return aMoney.amount * exchangeRates[key]
