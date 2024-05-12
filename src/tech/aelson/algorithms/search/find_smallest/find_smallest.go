package find_smallest

import "src/tech/aelson/m/v2/tech/aelson/algorithms/model"

func FindSmallest(products []*model.Product, start, end int) int {
	smallest := start
	for atual := start; atual <= end; atual++ {
		if products[atual].GetPrice() < products[smallest].GetPrice() {
			smallest = atual
		}
	}
	return smallest
}
