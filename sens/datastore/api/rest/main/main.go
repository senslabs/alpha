package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	SleepMain(r)
	DeviceMain(r)
	
	http.Handle("/", r)
}
