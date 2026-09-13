package functions

import (
	"fmt"
	"math/big"
	"uuid"

	"go_basics.2.0/data-access/data"
)

func SetOrganization() {
	_orgId := uuid.NewV7()

	organization := &data.Organization{
		Id:        _orgId,
		Name:      "Org_ABC",
		Employees: generateEmployees(_orgId),
	}

	printOrg(organization)
	printOrgEmployees(organization)

}

func generateEmployees(orgId uuid.UUID) data.Employees {
	_org := make(data.Employees)

	for i := range 5 {
		_uuid := uuid.NewV7()
		_experience := new(big.Int).Add(
			big.NewInt(int64(i)),
			new(big.Int).SetBytes(_uuid[:]),
		)
		_emp := data.Employee{
			Id:         _uuid,
			Name:       fmt.Sprintf("Employee_%d", i),
			Age:        i + 20,
			Experience: fmt.Sprintf("Experience_%d", _experience),
			Department: fmt.Sprintf("Department_%d", i),
			Position: data.Position{
				Id:   uuid.New(),
				Name: "Front End Engineer",
			},
		}
		_org[_uuid] = append(_org[_uuid], _emp)
	}
	return _org
}

func printOrg(org data.Logger) {
	org.DisplayOrgInfo()
}

func printOrgEmployees(employees data.Logger) {
	employees.DisplayOrgEmployees()
}
