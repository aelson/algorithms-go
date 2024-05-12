package find_smaller_elements

import "src/tech/aelson/m/v2/tech/aelson/algorithms/model"

func FindSmallerElements(reference model.Grade, unsortedGradle []model.Grade) int {
	lowerValuesCount := 0
	for _, grade := range unsortedGradle {
		if grade.GetResult() < reference.GetResult() {
			lowerValuesCount++
		}
	}
	return lowerValuesCount
}
