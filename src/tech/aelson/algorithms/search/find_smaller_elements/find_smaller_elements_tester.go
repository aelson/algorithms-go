package find_smaller_elements

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func FindSmallerElementsTester() {
	guilherme := model.Grade{StudentName: "guilherme", Result: 7}
	unsortedGrades := []model.Grade{
		{StudentName: "andre", Result: 4},
		{StudentName: "carlos", Result: 8.5},
		{StudentName: "ana", Result: 10},
		{StudentName: "jonas", Result: 3},
		{StudentName: "juliana", Result: 6.7},
		guilherme,
		{StudentName: "paulo", Result: 9},
		{StudentName: "mariana", Result: 5},
		{StudentName: "lucia", Result: 9.3},
	}

	lowerValuesCount := FindSmallerElements(guilherme, unsortedGrades)
	fmt.Println("Lower values count:", lowerValuesCount)
}
