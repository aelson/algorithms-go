package selection_sort

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/search"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func SelectionSort(products []*model.Product, numberOfElements int) {
	for current := 0; current < numberOfElements-1; current++ {
		fmt.Println("I am in the element", current)

		smallest := search.FindSmallest(products, current, len(products)-1)

		util.Swap(products, current, smallest)
	}
}
