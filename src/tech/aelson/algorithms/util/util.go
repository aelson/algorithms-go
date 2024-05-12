package util

import (
	"fmt"
	"src/tech/aelson/m/v2/tech/aelson/algorithms/model"
)

func SwapProduct(products []*model.Product, first, second int) {
	fmt.Println("<-> Swapping element", first, "with", second)
	fmt.Println("<-> Swapping product", products[first], "with", products[second])
	products[first], products[second] = products[second], products[first]
	fmt.Println("------------------------------------")
}

func SwapGrade(grades []*model.Grade, first int, second int) {
	fmt.Println("<-> Swapping element", first, "with", second)
	fmt.Println("<-> Swapping grade", grades[first], "with", grades[second])
	grades[first], grades[second] = grades[second], grades[first]
	fmt.Println("------------------------------------")
}
func PrintProductsArray(arrayTitle string, products []*model.Product) {
	fmt.Println(arrayTitle)
	for _, product := range products {
		fmt.Println(product.GetName(), "costs", product.GetPrice())
	}
}

func PrintGradesArray(arrayTitle string, grades []*model.Grade) {
	fmt.Println(arrayTitle)
	for _, grade := range grades {
		fmt.Println(grade.GetStudentName(), grade.GetResult())
	}
}

func GetUnsortedProducts() []*model.Product {
	return []*model.Product{
		model.NewProduct("Ford Escape", 30000),
		model.NewProduct("Toyota Corolla", 20000),
		model.NewProduct("Audi Q5", 45000),
		model.NewProduct("Honda Civic", 30000),
		model.NewProduct("Tesla Model 3", 50000),
	}
}

func GetUnsortedGrades(specialGrade *model.Grade) []*model.Grade {
	return []*model.Grade{
		model.NewGrade("andre", 4),
		model.NewGrade("carlos", 8.5),
		model.NewGrade("ana", 10),
		model.NewGrade("jonas", 3),
		model.NewGrade("juliana", 6.7),
		model.NewGrade("lucia", 9.3),
		model.NewGrade("paulo", 9),
		model.NewGrade("mariana", 5),
		specialGrade,
	}
}
