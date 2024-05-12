package model

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{Name: name, Price: price}
}

func (p *Product) String() string {
	return fmt.Sprintf("%s $%.2f", p.Name, p.Price)
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}
