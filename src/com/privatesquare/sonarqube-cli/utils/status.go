package utils

import "log"

/*
HandelGetRequest handles the status of a GET request made to service now
@param status response status of a request
return void
*/
func HandelGetStatus(status string) {
	if status == "" {
		log.Println("Success Ok")
	} else if status == "401 Unauthorized" {
		log.Fatal("Unauthorized, username or password is invalid")
	}
}
