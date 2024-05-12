package util

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func Swap(products []*model.Product, first, second int) {
	fmt.Println("<-> Swapping element", first, "with", second)

	firstProduct := products[first]
	secondProduct := products[second]

	fmt.Println("<-> Swapping product", firstProduct.GetName(), "with", secondProduct.GetName())

	products[first] = secondProduct
	products[second] = firstProduct

	fmt.Println("------------------------------------")
}

func PrintArray(arrayTitle string, products []*model.Product) {
	fmt.Println(arrayTitle)
	for _, product := range products {
		fmt.Println(product.GetName(), "costs", product.GetPrice())
	}
}

func GetUnsortedProducts() []*model.Product {
	return []*model.Product{
		model.NewProduct("Ford Escape", 30000),
		model.NewProduct("Toyota Corolla", 20000),
		model.NewProduct("Audi Q5", 45000),
		model.NewProduct("Honda Civic", 30000),
		model.NewProduct("Tesla Model 3", 50000),
	}
}
