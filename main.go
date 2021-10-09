package main

import (
	"log"
	"net/http"

	_ "nimeshjohari02.com/restapi/auth"
	_ "nimeshjohari02.com/restapi/instagram"
	_ "nimeshjohari02.com/restapi/user"
)

func main() {
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
