package main

import (
	"database/sql"
	"net/http"
)

//this function receive the sql.DB as input
//and use its Ping() method to check server validilty
//1st return http.StatusServiceUnavailable if server invalid
//1st return http.StatusOk if server valid
//2nd return false (unavailable) if server invalid
func dbAvailablityResponse(db *sql.DB) (ResponseBody, bool) {
	if db.Ping() != nil {
		return ResponseBody{
			Status: http.StatusServiceUnavailable,
			Message: "The server is currently unable to handle the request due to " +
				"a temporary overloading or maintenance of the server",
		}, false
	}

	return ResponseBody{
		Status:  http.StatusOK,
		Message: "Server Available",
	}, true
}
