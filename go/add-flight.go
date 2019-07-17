/********************************************************************
    file:   add-flight.go
    brief:  add a new flight info
********************************************************************/
package main

import "fmt"
import "net/http"
import "encoding/json"
import "strings"

import "./Database"

type AddFlightResponse struct {
	Result string
	Str    string
}

/********************************************************************
    func:   addFlight
	brief:  add a new flight info
	args:   w - responseWriter
			r - request
    return:
********************************************************************/
func addFlight(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------------------------------------Add Flight-------------------------------------")

	// get request
	r.ParseForm()
	form := r.Form
	date := form["date"][0]
	code := form["code"][0]

	// log
	fmt.Println("date: ", date, ", code: ", code)

	var response AddFlightResponse

	// search
	sql := "INSERT INTO `nyumeet`.`flight` (`code`, `date`) VALUES ('" + code + "', '" + date + "');"
	_, err := Database.Excute(sql)
	if err != nil {
		fmt.Println("error: ", err)
		if strings.HasPrefix (err.Error(), "Error 1062") {
			// Duplicate entry
			response.Result = "Fail"
			response.Str = "Duplicate entry"
		} else {
			response.Result = "Fail"
			response.Str = "Server error"
		}
	}

	if response.Result == "" {
		response.Result = "Success"
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
