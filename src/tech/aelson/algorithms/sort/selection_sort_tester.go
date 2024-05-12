package sort

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func SelectionSortTester() {
	products := []*model.Product{
		model.NewProduct("Ford Escape", 30000),
		model.NewProduct("Toyota Corolla", 20000),
		model.NewProduct("Audi Q5", 45000),
		model.NewProduct("Honda Civic", 30000),
		model.NewProduct("Tesla Model 3", 50000),
	}

	fmt.Println("Original array: ")
	for _, product := range products {
		fmt.Println(product.GetName(), "costs", product.GetPrice())
	}

	SelectionSort(products, len(products))

	fmt.Println("Sorted array: ")
	for _, product := range products {
		fmt.Println(product.GetName(), "costs", product.GetPrice())
	}
}
