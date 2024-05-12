package insertion_sort

import "src/tech/aelson/m/v2/tech/aelson/algorithms/util"

func InsertionSortTester() {
	products := util.GetUnsortedProducts()
	util.PrintProductsArray("Original array: ", products)

	InsertionSort(products, len(products))
	util.PrintProductsArray("Sorted array: ", products)
}
