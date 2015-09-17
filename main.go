package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	port := flag.Int("port", 5000, "port for listening")
	flag.Parse()

	r := mux.NewRouter()

	http.ListenAndServe(fmt.Sprintf(":%d", *port), r)
}
