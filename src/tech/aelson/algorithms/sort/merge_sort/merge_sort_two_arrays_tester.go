package merge_sort

import (
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func MergeSortTwoArraysTester() {
	firstArray := []*model.Grade{
		{StudentName: "andre", Result: 4},
		{StudentName: "mariana", Result: 5},
		{StudentName: "carlos", Result: 8.5},
		{StudentName: "paulo", Result: 9},
	}

	secondArray := []*model.Grade{
		{StudentName: "jonas", Result: 3},
		{StudentName: "juliana", Result: 6.7},
		{StudentName: "guilherme", Result: 7},
		{StudentName: "lucia", Result: 9.3},
		{StudentName: "ana", Result: 10},
	}

	util.PrintGradesArray("First array: ", firstArray)
	util.PrintGradesArray("Second array: ", secondArray)

	rank := MergeTwoArrays(firstArray, secondArray)
	util.PrintGradesArray("Merged array: ", rank)
}
