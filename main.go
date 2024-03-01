package main

import (
	"net/http"

	pauth "belajarpostman/jauth"
)

func main() {

	http.HandleFunc("/", pauth.Dindex)
	http.ListenAndServe(":3000", nil)

}
