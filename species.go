package swapi

import "fmt"

type Species struct {
	Name            string   `json:"name"`
	Classification  string   `json:"classification"`
	Designation     string   `json:"designation"`
	AverageHeight   string   `json:"average_height"`
	SkinColors      string   `json:"skin_colors"`
	HairColors      string   `json:"hair_colors"`
	EyeColors       string   `json:"eye_colors"`
	AverageLifespan string   `json:"average_lifespan"`
	Homeworld       string   `json:"homeworld"`
	Language        string   `json:"language"`
	People          []string `json:"people"`
	Films           []string `json:"films"`
	Created         string   `json:"created"`
	Edited          string   `json:"edited"`
	URL             string   `json:"url"`
}

func GetSpecies(id int) (Species, error) {
	var s Species
	return s, Get(fmt.Sprintf("/species/%d", id), &s)
}

func (s Species) GetFilms() (films []Film, err error) {
	return getFilms(s.Films)
}

func (s Species) GetPeople() (people []Person, err error) {
	return getPeople(s.People)
}
