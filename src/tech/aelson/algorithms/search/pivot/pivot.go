package pivot

import (
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/util"
)

func Pivot(grades []*model.Grade, end int) int {
	pivot := grades[end-1]
	lowerValuesCount := 0
	for index := 0; index < end-1; index++ {
		current := grades[index]
		if current.GetResult() < pivot.GetResult() {
			util.SwapGrade(grades, index, lowerValuesCount)
			lowerValuesCount++
		}
	}
	util.SwapGrade(grades, end-1, lowerValuesCount)
	return lowerValuesCount
}
