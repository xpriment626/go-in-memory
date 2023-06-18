package utils

import (
	"memory-crud/models"
)

func Populate() {
	models.Tenants = append(
		models.Tenants,
		models.Tenant{
			ID: "1", AreaCode: "0110",
			District: "Ur mom's basement",
			LandLord: &models.LandLord{FirstName: "John", LastName: "Doe"}})
}
