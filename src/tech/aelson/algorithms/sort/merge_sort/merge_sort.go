package merge_sort

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func MergeTwoArrays(firstArray, secondArray []*model.Grade) []*model.Grade {
	total := len(firstArray) + len(secondArray)
	merged := make([]*model.Grade, total)
	currentOfFirstArray := 0
	currentOfSecondArray := 0
	currentOfMerged := 0

	for currentOfFirstArray < len(firstArray) && currentOfSecondArray < len(secondArray) {
		grade1 := firstArray[currentOfFirstArray]
		grade2 := secondArray[currentOfSecondArray]

		fmt.Println("Comparing", grade1.GetStudentName(), "(", grade1.GetResult(), ") with", grade2.GetStudentName(), "(", grade2.GetResult(), ")")

		if grade1.GetResult() < grade2.GetResult() {
			fmt.Println("-> Inserting", grade1.GetStudentName(), "(", grade1.GetResult(), ") on the position", currentOfMerged)
			merged[currentOfMerged] = grade1
			currentOfFirstArray++
		} else {
			fmt.Println("-> Inserting", grade2.GetStudentName(), "(", grade2.GetResult(), ") on the position", currentOfMerged)
			merged[currentOfMerged] = grade2
			currentOfSecondArray++
		}
		fmt.Println("------------------------------------")
		currentOfMerged++
	}

	for currentOfFirstArray < len(firstArray) {
		fmt.Println("-> Inserting", firstArray[currentOfFirstArray].GetStudentName(), "(", firstArray[currentOfFirstArray].GetResult(), ") on the position", currentOfMerged, "because it is left over from the first array")
		merged[currentOfMerged] = firstArray[currentOfFirstArray]
		currentOfFirstArray++
		currentOfMerged++
	}

	for currentOfSecondArray < len(secondArray) {
		fmt.Println("-> Inserting", secondArray[currentOfSecondArray].GetStudentName(), "(", secondArray[currentOfSecondArray].GetResult(), ") on the position", currentOfMerged, "because it is left over from the second array")
		merged[currentOfMerged] = secondArray[currentOfSecondArray]
		currentOfSecondArray++
		currentOfMerged++
	}

	return merged
}

func SortOneArray(array []*model.Grade, start, middle, end int) []*model.Grade {
	total := len(array)
	sorted := make([]*model.Grade, total-start)
	sortedIndex := 0
	firstPartIndex := start
	secondPartIndex := middle

	for firstPartIndex < middle && secondPartIndex < end {
		fmt.Println("Comparing", array[firstPartIndex].GetStudentName(), "(", array[firstPartIndex].GetResult(), ") with", array[secondPartIndex].GetStudentName(), "(", array[secondPartIndex].GetResult(), ")")

		if array[firstPartIndex].GetResult() < array[secondPartIndex].GetResult() {
			fmt.Println("-> Inserting", array[firstPartIndex].GetStudentName(), "(", array[firstPartIndex].GetResult(), ") on the position", sortedIndex)
			sorted[sortedIndex] = array[firstPartIndex]
			firstPartIndex++
		} else {
			fmt.Println("-> Inserting", array[secondPartIndex].GetStudentName(), "(", array[secondPartIndex].GetResult(), ") on the position", sortedIndex)
			sorted[sortedIndex] = array[secondPartIndex]
			secondPartIndex++
		}
		fmt.Println("------------------------------------")
		sortedIndex++
	}

	for firstPartIndex < middle {
		fmt.Println("-> Inserting", array[firstPartIndex].GetStudentName(), "(", array[firstPartIndex].GetResult(), ") on the position", sortedIndex, "because it is left over from the first part of the array")
		sorted[sortedIndex] = array[firstPartIndex]
		firstPartIndex++
		sortedIndex++
	}

	for secondPartIndex < end {
		fmt.Println("-> Inserting", array[secondPartIndex].GetStudentName(), "(", array[secondPartIndex].GetResult(), ") on the position", sortedIndex, "because it is left over from the second part of the array")
		sorted[sortedIndex] = array[secondPartIndex]
		secondPartIndex++
		sortedIndex++
	}

	if start > 0 {
		fmt.Println("Rebuilding the original array keeping the initial object(s) not ordered (because the start is greater than 0)")
		for indexOfMerged := 0; indexOfMerged < sortedIndex; indexOfMerged++ {
			fmt.Println("-> Inserting", sorted[indexOfMerged].GetStudentName(), "(", sorted[indexOfMerged].GetResult(), ") on the position", start+indexOfMerged)
			array[start+indexOfMerged] = sorted[indexOfMerged]
		}
	}

	return array
}
