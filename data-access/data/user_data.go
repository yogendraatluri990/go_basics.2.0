package data

type Position struct {
	Id   string
	Name string
}
type Employee struct {
	Name       string
	Age        int
	Position   Position
	Experience string
	Id         string
	Department string
}

type Organization struct {
	Id        string
	Name      string
	Employees []Employee
}
