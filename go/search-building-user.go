/********************************************************************
    file:   search-building-user.go
    brief:  search users on a certain building
********************************************************************/
package main

import "fmt"
import "net/http"
import "encoding/json"

import "./Database"

// building info
type BuildingUserInfo struct {
	Name    string
	Room    string
	Wechat  string
	Message string
}

type SearchBuildingUserResponse struct {
	User []BuildingUserInfo
}

/********************************************************************
    func:   searchBuildingUser
	brief:  search users on a certain building
	args:   w - responseWriter
			r - request
    return:
********************************************************************/
func searchBuildingUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---------------------------------Search Building User----------------------------------")

	// get request
	r.ParseForm()
	form := r.Form
	building := form["building"][0]

	// log
	fmt.Println("building: ", building)

	var response SearchBuildingUserResponse

	// search
	sql := "SELECT DISTINCT * FROM `nyumeet`.`building-user` WHERE building='" + building + "';"
	rows, err := Database.Excute(sql)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for rows.Next() {
		var (
			building    string
			room    string
			wechat  string
			message string
		)
		rows.Scan(&building, &room, &wechat, &message)
		response.User = append(response.User, BuildingUserInfo{building, room, wechat, message})
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
