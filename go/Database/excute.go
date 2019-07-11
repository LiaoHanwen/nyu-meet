/********************************************************************
    file:   excute.go
    brief:  excute sql and return result
********************************************************************/
package Database

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

/********************************************************************
    func:   Excute
    brief:  excute sql and return result
	args:   sql
    return: results
********************************************************************/
func Excute(s string) (*sql.Rows, error) {
	// init database	
	db, err := sql.Open("mysql", <dataSourceName>)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// search
	rows, err := db.Query(s)

	return rows, err
}
