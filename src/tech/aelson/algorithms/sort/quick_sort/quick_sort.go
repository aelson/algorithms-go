package quick_sort

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/search/pivot"
)

func QuickSort(grades []*model.Grade, start, end int) {
	fmt.Println("Executing for", start, "-", end)
	numberOfElements := end - start
	if numberOfElements > 1 {
		pivotIndex := pivot.Pivot(grades, end)
		QuickSort(grades, start, pivotIndex)
		QuickSort(grades, pivotIndex+1, end)
	}
}
