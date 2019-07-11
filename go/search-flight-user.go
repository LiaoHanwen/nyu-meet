/********************************************************************
    file:   search-flight-user.go
    brief:  search users on a certain flight
********************************************************************/
package main

import "fmt"
import "net/http"
import "encoding/json"

import "./Database"

// flight info
type FlightUserInfo struct {
	Code    string
	Date    string
	Wechat  string
	Message string
}

type SearchFlightUserResponse struct {
	User []FlightUserInfo
}

/********************************************************************
    func:   searchFlightUser
	brief:  search users on a certain flight
	args:   w - responseWriter
			r - request
    return:
********************************************************************/
func searchFlightUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---------------------------------Search Flight User----------------------------------")

	// get request
	r.ParseForm()
	form := r.Form
	date := form["date"][0]
	code := form["code"][0]

	// log
	fmt.Println("date: ", date, ", code: ", code)

	var response SearchFlightUserResponse

	// search
	sql := "SELECT * FROM `nyumeet`.`flight-user` WHERE date='" + date + "'"
	rows, err := Database.Excute(sql)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for rows.Next() {
		var (
			code    string
			date    string
			wechat  string
			message string
		)
		rows.Scan(&code, &date, &wechat, &message)
		response.User = append(response.User, FlightUserInfo{code, date, wechat, message})
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
