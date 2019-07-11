/********************************************************************
    file:   main.go
    brief:  go web server
********************************************************************/

package main

import (
	"fmt"
	"log"
	"net/http"
)

/********************************************************************
    func:   defaultRoute
    brief:  return 404 and message
    args:   w - request ResponseWriter
            r - request
    return:
********************************************************************/
func defaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("---------------------------------Command Not Found-----------------------------------")
	http.Error(w, "go server command not found", 404)
	fmt.Println("-------------------------------------------------------------------------------------")
}

func main() {
	// set route
	http.HandleFunc("/", defaultRoute)
	http.HandleFunc("/go/search-flight", searchFlight)
	http.HandleFunc("/go/add-flight", addFlight)
	http.HandleFunc("/go/search-flight-user", searchFlightUser)
	http.HandleFunc("/go/add-flight-user", addFlightUser)

	// set listen port
	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
