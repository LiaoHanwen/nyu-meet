/********************************************************************
    file:   search-flight.go
    brief:  search flight on a certain day
********************************************************************/
package main

import "fmt"
import "net/http"
import "encoding/json"

import "./Database"

// flight info
type FlightInfo struct {
	Code string
	Date string
}

type SearchFlightResponse struct {
	Flight []FlightInfo
}

/********************************************************************
    func:   searchFlight
	brief:  search flight on a certain day
	args:   w - responseWriter
			r - request
    return:
********************************************************************/
func searchFlight(w http.ResponseWriter, r *http.Request) {
	fmt.Println("------------------------------------Search Flight------------------------------------")

	// get request
	r.ParseForm()
	form := r.Form
	date := form["date"][0]

	// log
	fmt.Println("date: ", date)

	var response SearchFlightResponse

	// search
	sql := "SELECT DISTINCT code, date FROM `nyumeet`.`flight-user` WHERE date='" + date + "'"
	rows, err := Database.Excute(sql)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for rows.Next() {
		var (
			code string
			date string
		)
		rows.Scan(&code, &date)
		response.Flight = append(response.Flight, FlightInfo{code, date})
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
