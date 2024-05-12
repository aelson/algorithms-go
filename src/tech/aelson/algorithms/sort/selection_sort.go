package sort

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/search"
)

func SelectionSort(products []*model.Product, numberOfElements int) {
	for current := 0; current < numberOfElements-1; current++ {
		fmt.Println("I am in the element", current)

		smallest := search.FindSmallest(products, current, len(products)-1)

		fmt.Println("<-> Swapping element", current, "with element", smallest)

		currentProduct := products[current]
		smallestProduct := products[smallest]

		fmt.Println("<-> Swapping product", currentProduct.GetName(), "with product", smallestProduct.GetName())

		products[current] = smallestProduct
		products[smallest] = currentProduct
		fmt.Println("------------------------------------")
	}
}
