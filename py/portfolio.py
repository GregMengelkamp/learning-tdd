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
        total = functools.reduce(operator.add, map(lambda m: m.amount, self.moneys), 0)
        return Money(total, currency)
