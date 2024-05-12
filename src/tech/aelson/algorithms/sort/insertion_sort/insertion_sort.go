package insertion_sort

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func InsertionSort(products []*model.Product, numberOfElements int) {
	for current := 1; current < numberOfElements; current++ {
		fmt.Println("I am in the element", current)
		elementBeingAnalysed := current
		for elementBeingAnalysed > 0 && products[elementBeingAnalysed].GetPrice() < products[elementBeingAnalysed-1].GetPrice() {
			util.SwapProduct(products, elementBeingAnalysed, elementBeingAnalysed-1)
			elementBeingAnalysed--
		}
	}
}
