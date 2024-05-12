package merge_sort

import (
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func MergeSortOneArrayTester() {
	grades := []*model.Grade{
		{StudentName: "andre", Result: 4},
		{StudentName: "mariana", Result: 5},
		{StudentName: "carlos", Result: 8.5},
		{StudentName: "paulo", Result: 9},
		{StudentName: "jonas", Result: 3},
		{StudentName: "juliana", Result: 6.7},
		{StudentName: "guilherme", Result: 7},
		{StudentName: "lucia", Result: 9.3},
		{StudentName: "ana", Result: 10},
	}
	util.PrintGradesArray("Grades array: ", grades)

	rank := SortOneArray(grades, 0, 4, len(grades))
	util.PrintGradesArray("Sorted array: ", rank)
}
