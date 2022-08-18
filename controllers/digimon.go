package controllers

import "example.com/m/api"

// type DIGIMON struct {
// 	Name string
// 	Id   string
// }

func GetDigiByName(name string) api.DIGIMONDATA {

	// d := DIGIMON{Name: name}
	return api.GetDigi(name)
}
