package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	AuthMain(r)
	OrgMain(r)
	OpMain(r)
	UserMain(r)
	EndpointMain(r)
	OrgAuthMain(r)
	OpAuthMain(r)
	UserAuthMain(r)
	OrgOpMain(r)
	OrgUserMain(r)
	OpUserMain(r)
	OrgEndpointMain(r)
	OpEndpointMain(r)
	UserEndpointMain(r)
	
	http.Handle("/", r)
}
