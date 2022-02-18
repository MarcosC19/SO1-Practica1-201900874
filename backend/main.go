package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	port     = 27017
	user     = "mongoadmin"
	password = "practica1-so"
)

type operation struct { // OPERATION JSON STRUCT
	Num1   float32   `json:"num1"`
	Num2   float32   `json:"num2"`
	Option string    `json:"operation"`
	Result float32   `json:"result"`
	Date   time.Time `json:"date"`
}

type modelMongo struct { // MODEL STRUCT COLLECTION WITHOUT ID
	ID        string    `bson:"-"`
	Number1   float32   `bson:"number1"`
	Number2   float32   `bson:"number2"`
	Operation string    `bson:"operation"`
	Result    float32   `bson:"result"`
	Date      time.Time `bson:"date"`
}

type modelMongoR struct { // MODEL STRUCT COLLECTION WITH ID
	ID        string    `bson:"_id"`
	Number1   float32   `bson:"number1"`
	Number2   float32   `bson:"number2"`
	Operation string    `bson:"operation"`
	Result    float32   `bson:"result"`
	Date      time.Time `bson:"date"`
}

type listOperation struct { // DATA RETURN COLLECTION
	Data []modelMongoR `json:"data"`
}

func saveOperation(newOperation operation) {
	host, defined := os.LookupEnv("HOSTIP")
	if !defined {
		host = "localhost"
	}

	// OPENING CONNECTION TO MONGODB
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// VERIFYING THE CONNECTION
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion exitosa")

	// CREATING A MODEL COLLECTION
	newOperationMongoDB := modelMongo{
		Number1:   newOperation.Num1,
		Number2:   newOperation.Num2,
		Operation: newOperation.Option,
		Result:    newOperation.Result,
		Date:      newOperation.Date,
	}

	// CONNECTION TO DATABASE AND COLLECTION
	collection := client.Database("practica1-so").Collection("operations")

	// INSERT THE NEW OPERATION
	insertResult, err := collection.InsertOne(context.TODO(), newOperationMongoDB)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nueva operacion insertada con exito ", insertResult)

	//CLOSING CONNECTION TO MONGODB
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion cerrada")
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the go server!")
}

func doOperation(w http.ResponseWriter, r *http.Request) {
	var newOperation operation // VARIABLE TO CONTAIN THE NEW OPERATION
	// READING BODY REQUEST
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newOperation)

	// SET HEADERS
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// READING TYPE OF OPERATION
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

	// SAVING THE NEW OPERATION
	saveOperation(newOperation)

	// RETURNING THE REQUEST
	json.NewEncoder(w).Encode(newOperation)
}

func getOperations(w http.ResponseWriter, r *http.Request) {
	host, defined := os.LookupEnv("HOSTIP")
	if !defined {
		host = "localhost"
	}

	// OPENING CONNECTION TO MONGODB
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", user, password, host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)

	if err != nil {
		log.Fatal(err)
	}

	// VERIFYING THE CONNECTION
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexion exitosa")

	// SET FIND OPTIONS
	findOptions := options.Find()

	// LIST OF OPERATION RESULTS
	var results []modelMongoR

	// CONNECTION TO DATABASE AND COLLECTION
	collection := client.Database("practica1-so").Collection("operations")

	// GET OPERATIONS
	current, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		log.Fatal(err)
	}

	// DECODING THE RESULTS
	for current.Next(context.TODO()) {

		var n modelMongoR
		err := current.Decode(&n)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, n)
	}

	if err := current.Err(); err != nil {
		log.Fatal(err)
	}

	current.Close(context.TODO())

	// SET JSON RETURN
	var retorno = listOperation{
		Data: results,
	}

	// SET HEADERS
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// RETURNING THE REQUEST
	json.NewEncoder(w).Encode(retorno)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)                                      // MAIN ROUTE
	router.HandleFunc("/Operation", doOperation).Methods("POST")      // OPERATION ROUTE
	router.HandleFunc("/getOperations", getOperations).Methods("GET") // OPERATIONS ROUTE
	fmt.Println("Server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
