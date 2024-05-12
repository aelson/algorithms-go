package pivot

import "src/tech/aelson/m/v2/tech/aelson/algorithms/model"

func Pivot(grades []model.Grade, start, end int) int {
	pivot := grades[end-1]
	lowerValuesCount := 0
	for index := 0; index < end-1; index++ {
		current := grades[index]
		if current.GetResult() <= pivot.GetResult() {
			swap(grades, index, lowerValuesCount)
			lowerValuesCount++
		}
	}
	swap(grades, end-1, lowerValuesCount)
	return lowerValuesCount
}

func swap(grades []model.Grade, from, to int) {
	grade1 := grades[from]
	grade2 := grades[to]
	grades[to] = grade1
	grades[from] = grade2
}
