package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type operation struct {
	Num1   float32 `json:"num1"`
	Num2   float32 `json:"num2"`
	Option string  `json:"operation"`
	Result float32 `json:"result"`
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the go server!")
}

func getOperation(w http.ResponseWriter, r *http.Request) {
	var newOperation operation

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newOperation)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	if newOperation.Option == "+" {
		newOperation.Result = newOperation.Num1 + newOperation.Num2
	} else if newOperation.Option == "-" {
		newOperation.Result = newOperation.Num1 - newOperation.Num2
	} else if newOperation.Option == "*" {
		newOperation.Result = newOperation.Num1 * newOperation.Num2
	} else if newOperation.Option == "/" {
		newOperation.Result = newOperation.Num1 / newOperation.Num2
	} else {
		newOperation.Result = 0
	}

	json.NewEncoder(w).Encode(newOperation)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/Operation", getOperation)
	fmt.Println("Server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
