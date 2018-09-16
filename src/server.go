package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//KvJSON is the structure for the response
type KvJSON struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
}

func kvServiceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Calling kvService")
	vars := mux.Vars(r)
	k := vars["key"]
	kv := KvJSON{}
	_, ok := KeyValMap[k]
	if ok {
		kv.Key = k
		kv.Value = KeyValMap[k]
		fmt.Println(kv)
		json.NewEncoder(w).Encode(kv)
	} else {
		fmt.Println("The requested key does not exist in the CSV")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("The requested key does not exist in the CSV\n"))
	}
}

func server(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/app/v1/kv/{key}", kvServiceHandler).Methods("GET")
	log.Printf("kvService running on: %s", addr)
	http.ListenAndServe(addr, router)
}
