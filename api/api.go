package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type DIGIMONDATA struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	XAntibody bool   `json:"xAntibody"`
	Images    []struct {
		Href        string `json:"href"`
		Transparent bool   `json:"transparent"`
	} `json:"images"`
	Levels []struct {
		ID    int    `json:"id"`
		Level string `json:"level"`
	} `json:"levels"`
	Types []struct {
		ID   int    `json:"id"`
		Type string `json:"type"`
	} `json:"types"`
	Attributes []struct {
		ID        int    `json:"id"`
		Attribute string `json:"attribute"`
	} `json:"attributes"`
	Fields []struct {
		ID    int    `json:"id"`
		Field string `json:"field"`
	} `json:"fields"`
	Descriptions []struct {
		Origin      string `json:"origin"`
		Language    string `json:"language"`
		Description string `json:"description"`
	} `json:"descriptions"`
}

func GetDigi(digimon string) DIGIMONDATA {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://digi-api.com/api/v1/digimon/%s", digimon), nil)
	if err != nil {
		//return err
		fmt.Print("error")
	}
	resp, err := client.Do(req)
	if err != nil {
		//return  err
		fmt.Print(err)
	}
	decoder := json.NewDecoder(resp.Body)
	var data DIGIMONDATA
	err = decoder.Decode(&data)
	if err != nil {
		//return err
		fmt.Print("error")
	}
	fmt.Print(data)
	return data
}

// func main() {
// 	digimon := getDigi("agumon")
// 	fmt.Println(digimon.Name)
// }
//https://www.youtube.com/watch?v=QP0slFxVQKU
