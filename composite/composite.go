package main

import (
	"fmt"
)

func printPrice(name string, price int) {
	fmt.Printf("%s: %d\n", name, price)
}

// Component type declare component interface.
type Component interface {
	GetPrice() int
}

// Product type declare Leaf structure.
type Product struct {
	Name  string
	Price int
}

// GetPrice implement Component interface for Leaf.
func (p Product) GetPrice() int {
	printPrice(p.Name, p.Price)
	return p.Price
}

// CompositeProduct type declare Branch structure.
type CompositeProduct struct {
	Name     string
	Children []Component
}

// GetPrice implement Component interface for Branch.
func (c CompositeProduct) GetPrice() int {
	price := 0
	for _, p := range c.Children {
		price += p.GetPrice()
	}
	printPrice(c.Name, price)
	return price
}

func main() {
	turbocharger := Product{"turbocharger", 500}
	intercooler := Product{"intercooler", 200}
	engineBlock := Product{"engineBlock", 1000}
	engine := CompositeProduct{"engine", []Component{turbocharger, intercooler, engineBlock}}

	tire := Product{"tire", 50}
	rim := Product{"rim", 100}
	wheel := CompositeProduct{"wheel", []Component{tire, rim}}

	gearbox := Product{"gearbox", 500}

	body := Product{"body", 1500}

	car := CompositeProduct{"car", []Component{body, engine, gearbox, wheel, wheel, wheel, wheel}}

	car.GetPrice()
}
