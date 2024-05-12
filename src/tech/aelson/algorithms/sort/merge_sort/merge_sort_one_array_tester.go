package merge_sort

import (
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func MergeSortOneArrayTester() {
	//sortedHalfsOfGrades := []*model.Grade{
	//	{StudentName: "mariana", Result: 5},
	//	{StudentName: "carlos", Result: 8.5},
	//	{StudentName: "lucia", Result: 9.3},
	//	{StudentName: "ana", Result: 10},
	//
	//	{StudentName: "jonas", Result: 3},
	//	{StudentName: "andre", Result: 4},
	//	{StudentName: "juliana", Result: 6.7},
	//	{StudentName: "guilherme", Result: 7},
	//	{StudentName: "paulo", Result: 9},
	//}
	unsortedGrades := []*model.Grade{
		{StudentName: "mariana", Result: 5},
		{StudentName: "ana", Result: 10},
		{StudentName: "carlos", Result: 8.5},
		{StudentName: "lucia", Result: 9.3},
		{StudentName: "andre", Result: 4},
		{StudentName: "paulo", Result: 9},
		{StudentName: "juliana", Result: 6.7},
		{StudentName: "jonas", Result: 3},
		{StudentName: "guilherme", Result: 7},
	}
	//util.PrintGradesArray("Grades array: ", sortedHalfsOfGrades)
	util.PrintGradesArray("Grades array: ", unsortedGrades)

	//SortOneArrayWithTwoOrderedHalfs(sortedHalfsOfGrades, 0, 4, len(sortedHalfsOfGrades))
	MergeSort(unsortedGrades, 0, len(unsortedGrades))
	util.PrintGradesArray("Sorted array: ", unsortedGrades)
}
