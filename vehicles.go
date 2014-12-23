package swapi

import "fmt"

type Vehicle struct {
	Name                 string
	Model                string
	Manufacturer         string
	CostInCredits        string `json:"cost_in_credits"`
	Length               string
	MaxAtmospheringSpeed string `json:"max_atmosphering_speed"`
	Crew                 string
	Passengers           string
	CargoCapacity        string `json:"cargo_capacity"`
	Consumables          string
	VehicleClass         string `json:"vechicle_class"`
	Pilots               []string
	Films                []string
	Created              string
	Edited               string
	URL                  string `json:"url"`
}

func GetVehicle(id int) (Vehicle, error) {
	var v Vehicle
	return v, Get(fmt.Sprintf("/vehicles/%d", id), &v)
}
