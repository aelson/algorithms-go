package selection_sort

import (
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func SelectionSortTester() {
	products := util.GetUnsortedProducts()
	util.PrintProductsArray("Original array: ", products)

	SelectionSort(products, len(products))
	util.PrintProductsArray("Sorted array: ", products)
}
