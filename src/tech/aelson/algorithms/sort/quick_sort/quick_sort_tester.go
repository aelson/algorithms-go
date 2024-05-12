package quick_sort

import (
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func QuickSortTester() {
	guilherme := model.NewGrade("guilherme", 7)
	unsortedGrades := util.GetUnsortedGrades(guilherme)

	QuickSort(unsortedGrades, 0, len(unsortedGrades))
	util.PrintGradesArray("Sorted array: ", unsortedGrades)
}
