package find_smallest

import "src/tech/aelson/m/v2/tech/aelson/algorithms/model"

func FindSmallest(products []*model.Product, start, end int) int {
	smallest := start
	for index := start; index <= end; index++ {
		if products[index].GetPrice() < products[smallest].GetPrice() {
			smallest = index
		}
	}
	return smallest
}
