package merge_sort

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func MergeSort(firstArray, secondArray []*model.Grade) []*model.Grade {
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
