/********************************************************************
    file:   search-building.go
    brief:  search building list
********************************************************************/
package main

import "fmt"
import "net/http"
import "encoding/json"

import "./Database"

// building info
type BuildingInfo struct {
	Name string
}

type SearchBuildingResponse struct {
	Building []BuildingInfo
}

/********************************************************************
    func:   searchBuilding
	brief:  search building list
	args:   w - responseWriter
			r - request
    return:
********************************************************************/
func searchBuilding(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------------------------Search Building------------------------------------")

	var response SearchBuildingResponse

	// search
	sql := "SELECT DISTINCT building FROM `nyumeet`.`building-user`;"
	rows, err := Database.Excute(sql)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for rows.Next() {
		var (
			building string
		)
		rows.Scan(&building)
		response.Building = append(response.Building, BuildingInfo{building})
	}

	// output
	jsonByte, err := json.Marshal(response)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Fprintf(w, string(jsonByte))

	fmt.Println("Return:", string(jsonByte))
	fmt.Println("-------------------------------------------------------------------------------------")
}
