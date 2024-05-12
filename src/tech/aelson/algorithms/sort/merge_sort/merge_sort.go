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

	currentOfMerged = addRemainingElementsToEndOfArray(firstArray, len(firstArray), currentOfFirstArray, merged, currentOfMerged)
	addRemainingElementsToEndOfArray(secondArray, len(secondArray), currentOfSecondArray, merged, currentOfMerged)

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

	sortedIndex = addRemainingElementsToEndOfArray(array, middle, firstPartIndex, sorted, sortedIndex)
	addRemainingElementsToEndOfArray(array, end, secondPartIndex, sorted, sortedIndex)
	if start+end < len(array) {
		rebuildArray(array, start, sortedIndex, sorted)
	}

	return array
}

func addRemainingElementsToEndOfArray(array []*model.Grade, arrayEnd, arrayIndex int, merged []*model.Grade, mergedArrayIndex int) int {
	for arrayIndex < arrayEnd {
		fmt.Println("-> Inserting", array[arrayIndex].GetStudentName(), "(", array[arrayIndex].GetResult(), ") on the position", mergedArrayIndex, "because it is left over from the first array")
		merged[mergedArrayIndex] = array[arrayIndex]
		arrayIndex++
		mergedArrayIndex++
	}
	return mergedArrayIndex
}

func rebuildArray(array []*model.Grade, start, sortedIndex int, sorted []*model.Grade) {
	fmt.Println("Rebuilding the original array")
	for indexOfMerged := 0; indexOfMerged < sortedIndex; indexOfMerged++ {
		fmt.Println("-> Inserting", sorted[indexOfMerged].GetStudentName(), "(", sorted[indexOfMerged].GetResult(), ") on the position", start+indexOfMerged)
		array[start+indexOfMerged] = sorted[indexOfMerged]
	}
}
