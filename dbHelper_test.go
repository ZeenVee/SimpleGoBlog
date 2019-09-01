package main

//unknown db input
//test fail during db.Ping()

/*func TestDBAvailablityResponse(t *testing.T) {
	tables := []struct {
		inDB            *sql.DB
		outStatus       int
		outMessage      string
		outAvailability bool
	}{
		{&sql.DB{},
			http.StatusServiceUnavailable,
			"The server is currently unable to handle the request due to " +
				"a temporary overloading or maintenance of the server",
			false,
		},
	}

	for _, table := range tables {
		resultBody, resultAvailability := dbAvailablityResponse(table.inDB)

		if resultBody.Status != table.outStatus {
			t.Errorf("dbAvailablityResponse(%v) Status was incorrect, got %v, want %v.",
				table.inDB, resultBody.Status, table.outStatus)
		}

		if resultBody.Message != table.outMessage {
			t.Errorf("dbAvailablityResponse(%v) Message was incorrect, got %v, want %v.",
				table.inDB, resultBody.Message, table.outMessage)
		}

		if resultAvailability != table.outAvailability {
			t.Errorf("dbAvailablityResponse(%v) validity was incorrect, got %v, want %v.",
				table.inDB, resultAvailability, table.outAvailability)
		}
	}
}*/
