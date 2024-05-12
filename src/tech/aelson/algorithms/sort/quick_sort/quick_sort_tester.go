package quick_sort

import (
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func QuickSortTester() {
	guilherme := model.NewGrade("guilherme", 7)
	unsortedGrades := []*model.Grade{
		{StudentName: "andre", Result: 4},
		{StudentName: "carlos", Result: 8.5},
		{StudentName: "ana", Result: 10},
		{StudentName: "jonas", Result: 3},
		{StudentName: "juliana", Result: 6.7},
		{StudentName: "lucia", Result: 9.3},
		{StudentName: "paulo", Result: 9},
		{StudentName: "mariana", Result: 5},
		guilherme,
	}

	QuickSort(unsortedGrades, 0, len(unsortedGrades))
	util.PrintGradesArray("Sorted array: ", unsortedGrades)
}
