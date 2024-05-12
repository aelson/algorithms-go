package pivot

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func PivotTester() {
	guilherme := model.Grade{StudentName: "guilherme", Result: 7}
	unsortedGrades := []model.Grade{
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

	pivotPosition := Pivot(unsortedGrades, 0, len(unsortedGrades))
	fmt.Println("The Pivot is in the position", pivotPosition)
	for _, grade := range unsortedGrades {
		fmt.Println(grade.GetStudentName(), grade.GetResult())
	}
}
