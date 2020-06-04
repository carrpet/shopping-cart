package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	port = "8080"
)

type frontendserver struct {
}

func main() {
	srvPort := port
	if os.Getenv("PORT") != "" {
		srvPort = os.Getenv("PORT")
	}
	addr := os.Getenv("LISTEN_ADDR")
	r := mux.NewRouter()
	//server := new(frontendserver)
	r.HandleFunc("/", homeHandler).Methods(http.MethodGet, http.MethodHead)
	fmt.Println("starting server on " + addr + ":" + srvPort)
	var handler http.Handler = r
	http.ListenAndServe(addr+":"+srvPort, handler)
}
