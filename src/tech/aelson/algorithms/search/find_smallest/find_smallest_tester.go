package find_smallest

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func TestSmallestPrice() {
	products := []*model.Product{
		model.NewProduct("Tesla Model 3", 50000),
		model.NewProduct("Toyota Corolla", 20000),
		model.NewProduct("Ford Escape", 30000),
		model.NewProduct("Honda Civic", 30000),
		model.NewProduct("Audi Q5", 45000),
	}

	smallest := FindSmallest(products, 0, 4)
	fmt.Println(smallest)
	fmt.Println("The car", products[smallest].GetName(), "costs", products[smallest].GetPrice())
}
