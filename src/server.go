package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func kvServiceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling kvService")
}

func server(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/app/v1/kv", kvServiceHandler).Methods("GET")
	log.Printf("kvService running on: %s", addr)
	http.ListenAndServe(addr, router)
}
