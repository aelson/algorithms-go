package model

import "fmt"

type Grade struct {
	StudentName string
	Result      float64
}

func NewGrade(studentName string, result float64) *Grade {
	return &Grade{StudentName: studentName, Result: result}
}

func (g *Grade) String() string {
	return fmt.Sprintf("%s: %d", g.StudentName, g.Result)
}

func (g *Grade) GetStudentName() string {
	return g.StudentName
}

func (g *Grade) GetResult() float64 {
	return g.Result
}
