### Learning TDD by applying it and hopefully acquiring some Go on the fly

imho, a developer should have three things above all:

- creativity
- blocks to build with
- know-how (to build stuff)

I consider myself to be reasonably creative.

By now I think I have enough experience in a programming language to say that I am in possession of solid, different elements to be able to build a lot of things with it.

So at the moment I'm mainly concerned with the "how?" and that's where I came across the book:
**"Learning Test-Driven Development -A Polyglot Guide to Writing Uncluttered Code"**
_by Saleem Siddiqui_

The blurb appealed to me, as did the structure after a first skim. I also like to take the opportunity to get to know "Go" and to expand and consolidate my knowledge of JavaScript.

### Below I will provide a summary of the project for my own review if you're interested.

#### 1st Problem: building a multi-currency spreadsheet to track a stock portfolio for example

##### 1st sub-problem: first feature to be implemented 5 USD × 2 = 10 USD

- write failing test
- implement what's necessary
- refactor: remove (logical coupling) _which is often root of duplicated code_ and tidy things up

##### 2nd sub-problem: abstract a **dollar** to genereic money which encapsules a value and a currency

- write failing test
- implement what's necessary
- refactor: remove (logical coupling) _which is often root of duplicated code_ and tidy things up
  - realization: money abstracts dollar. refactor dollar due to DRY principle

##### 3rd sub-problem: allow division

- you now know the drill...
- observations: Go is so not dynamically typed Qq..., new way of testing by comparing whole structs/objects (actual and expected)
- refactor: helper function in Go, deepStrictEqual and inlining call of arithmetic method in JS, overwrite of eq-method to compare objects directly not only by reference but also by value

##### 4th sub problem: addition of different currencies

- _5 USD + 10 EUR = 17 USD adding dollars to dollars results in dollars_; exploring problem **domain** to discover new
  entities, relationships, functions, and methods; outcome: collection of currencies (portfolio) that should be able to evaluate;
  a test to add two Money entities in the same currency, using the Portfolio as a new entity
- you define how fast you want to go; start with trivial implementations; adding the same currencies first; mixing currencies would require exchange rates which are not conceptualized yet
-
- observations: answering the question **test-drive or use “refactoring budget” (while you have green tests) to do it?**; forces "design thinking" and determines speed

##### recap // new stuff

- logical coupling is often times equivalent to duplicated code - DRY principle
- eq-method in python checks equality by reference by default
- explore a problem _domain_ to come up with new entities, relationsships, functions and methods
- don't fear to start with trivial implementations
- Dependency injection: separate creation of an object from usage -> increases cohesion, reduces coupling
- you can be pretty functional with python
- sometimes statically typed languages are not completely awful

- separate source files, still one package; package management in Go
- Reflect is a very handy and interesting JS object
- circle: Where do you get the data you need to achieve what you want to do? (database, API, **Prototype!?**)

##### 5th sub problem: conversion between currencies

- some onsiderations concerning the problem domain: pair of currencies, definite **UNI**directional exchange rate, not always arithmetically reciprocal, currencies may be inconvertible
- red: implement features sequentially by starting with a trivial approach if necessary; "destructure" and expand
- green: added a working solution for converting EUR to USD; conversion between currencies with predefined values added, errorhandling for missing exchange rates
- refactor:

- newly learned about the _transformation priority premise_, "comma, ok" idiom in Go, structs can implement one or more interfaces in Go, lambda can get messy quickly in python, no simple way to catch exceptions with lambds,

##### 6th sub problem: unclutter code, avoid dependecies and duplications

- red: domain driven design (ddd) related a bank entity makes sense to take care of the currency exchange. it shouldn't be a responsibility associated with a portfolio; writing tests for conversion with a bank
- green: implementing and injecting _bank_ to the portfolio.evaluate, using the key to directly form errormessages
- refactor: deleting the former portfolio.conversion functions, updating the tests to reflect changes to portfolio.evaluate requiring a bank
