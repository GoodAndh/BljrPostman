package main

import (
	"net/http"

	pauth "belajarpostman/jauth"
)

func main() {

	http.HandleFunc("/", pauth.Index)
	http.ListenAndServe(":3000", nil)

}
