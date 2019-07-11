/********************************************************************
    file:   add-flight-user.go
    brief:  add a new flight user
********************************************************************/
package main

import "fmt"
import "net/http"
import "encoding/json"

import "./Database"

type AddFlightUserResponse struct {
	Result string
	Str    string
}

/********************************************************************
    func:   addFlightUser
	brief:  add a new flight user
	args:   w - responseWriter
			r - request
    return:
********************************************************************/
func addFlightUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-----------------------------------Add Flight User-----------------------------------")

	// get request
	r.ParseForm()
	form := r.Form
	date := form["date"][0]
	code := form["code"][0]
	wechat := form["wechat"][0]
	message := form["message"][0]

	// log
	fmt.Println("date: ", date, ", code: ", code, ", wechat: ", wechat)
	fmt.Println("message: ", message)

	var response AddFlightUserResponse
	result := true

	// search
	sql := "INSERT INTO `nyumeet`.`flight-user` (`code`, `date`, `wechat`, `message`) VALUES ('" + code + "', '" + date + "', '" + wechat + "', '" + message + "');"
	_, err := Database.Excute(sql)
	if err != nil {
		fmt.Println("error: ", err)
		result = false
	}

	if result {
		response.Result = "Success"
	} else {
		response.Result = "Fail"
		response.Str = "Server error"
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
