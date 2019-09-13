from abc import ABCMeta, abstractmethod
import inspect


class Component(metaclass=ABCMeta):
    @abstractmethod
    def get_price(self):
        pass

    def print(self, price, depth):
        indent = ' ' * (depth - 2) * 2
        print('{}{}: {}'.format(indent, self.name, price))


class Product(Component):
    def __init__(self, name, price):
        self.name = name
        self.price = price
    
    def get_price(self):
        self.print(self.price, len(inspect.stack()))
        return self.price


class CompositeProduct(Component):
    def __init__(self, name, children=list()):
        self.name = name
        self.children = children

    def get_price(self):
        price = sum(child.get_price() for child in self.children)
        self.print(price, len(inspect.stack()))
        return price


turbocharger = Product('turbocharger', 500)
intercooler = Product('intercooler', 200)
engine_block = Product('engine_block', 1000)
engine = CompositeProduct('engine', [turbocharger, intercooler, engine_block])

tire = Product('tire', 50)
rim = Product('rim', 100)
wheel = CompositeProduct('wheel', [tire, rim])

gearbox = Product('gearbox', 500)

body = Product('body', 1500)

car = CompositeProduct('car', [body, engine, gearbox, wheel, wheel, wheel, wheel])

car.get_price()
