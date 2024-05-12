package model

type Grade struct {
	StudentName string
	Result      float64
}

func NewGrade(studentName string, result float64) *Grade {
	return &Grade{StudentName: studentName, Result: result}
}

func (p *Grade) GetStudentName() string {
	return p.StudentName
}

func (p *Grade) GetResult() float64 {
	return p.Result
}
