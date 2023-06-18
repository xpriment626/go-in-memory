package models

type Tenant struct {
	ID       string    `json:"id"`
	AreaCode string    `json:"property_id"`
	District string    `json:"district"`
	LandLord *LandLord `json:"landlord"`
}

type LandLord struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var Tenants []Tenant
