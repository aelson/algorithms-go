package insertion_sort

import "src/tech/aelson/m/v2/tech/aelson/algorithms/util"

func InsertionSortTester() {
	products := util.GetUnsortedProducts()
	util.PrintArray("Original array: ", products)

	InsertionSort(products, len(products))
	util.PrintArray("Sorted array: ", products)
}
