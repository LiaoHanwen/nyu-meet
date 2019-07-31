/********************************************************************
    file:   search-flight-time.go
    brief:  search flight user by arrive time
********************************************************************/
package main

import "fmt"
import "net/http"
import "strings"
import "encoding/json"

import "./Database"

// flight info
type FlightUserFullInfo struct {
	Code string
	Date string
	Wechat string
	Dep	string
	Arr string
	DepTime string
	ArrTime string
	Message string
}

type SearchFlightTimeResponse struct {
	User []FlightUserFullInfo
}

/********************************************************************
    func:   searchFlightTime
	brief:  search flight user by arrive time
	args:   w - responseWriter
			r - request
    return:
********************************************************************/
func searchFlightTime(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---------------------------------Search Flight Time----------------------------------")

	// get request
	r.ParseForm()
	form := r.Form
	date := form["date"][0]
	airport := form["airport"][0]
	start := form["start"][0]
	end := form["end"][0]

	// log
	fmt.Println("date: ", date)

	var response SearchFlightTimeResponse

	// search
	sql := "SELECT DISTINCT `nyumeet`.`flight`.code, departure, arrive, deptime, arrtime, date, wechat, message FROM `nyumeet`.`flight` INNER JOIN `nyumeet`.`flight-user` ON  `nyumeet`.`flight`.`code` = `nyumeet`.`flight-user`.`code` where date='" + date + "' and arrtime<'" + end + "' and arrtime>'" + start + "'"
	if airport != "" {
		sql = sql + " and arrive='" + airport + "';"
	} else {
		sql += ";"
	}

	fmt.Println(sql)

	rows, err := Database.Excute(sql)
	if err != nil {
		fmt.Println("error: ", err)
	}

	for rows.Next() {
		var (
			code string
			date string
			departure string
			arrive string
			deptime string
			arrtime string
			wechat string
			message string
		)
		rows.Scan(&code, &departure, &arrive, &deptime, &arrtime, &date, &wechat, &message)
		deptime = strings.TrimSuffix(deptime, ":00")
		arrtime = strings.TrimSuffix(arrtime, ":00")
		response.User = append(response.User, FlightUserFullInfo{code, date, wechat, departure, arrive, deptime, arrtime, message})
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
