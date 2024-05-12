package pivot

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func PivotTester() {
	guilherme := model.NewGrade("guilherme", 7)
	unsortedGrades := util.GetUnsortedGrades(guilherme)

	pivotPosition := Pivot(unsortedGrades, len(unsortedGrades))
	fmt.Println("The Pivot is in the position", pivotPosition)
	for _, grade := range unsortedGrades {
		fmt.Println(grade.GetStudentName(), grade.GetResult())
	}
}
