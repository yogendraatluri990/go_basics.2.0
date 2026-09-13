package data

import (
	"fmt"
	"uuid"
)

type Position struct {
	Id   uuid.UUID
	Name string
}
type Employee struct {
	Name       string
	Age        int
	Position   Position
	Experience string
	Id         uuid.UUID
	Department string
}

type Employees map[uuid.UUID][]Employee
type Organization struct {
	Id        uuid.UUID
	Name      string
	Employees Employees
}

func (org *Organization) DisplayOrgInfo() {
	fmt.Printf("\t %+v\n", *org)
}

func (org *Organization) DisplayOrgEmployees() {
	fmt.Println("Employees Len:", len(org.Employees))
	for _, employee := range org.Employees {
		fmt.Printf("\t %+v\n", employee)
	}
}

type Logger interface {
	DisplayOrgInfo()
	DisplayOrgEmployees()
}
