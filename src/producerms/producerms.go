package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

// MongoDB Config
//var mongodb_server = "mongodb"
var mongodb_server = "mongo"
var mongodb_database = "cmpe281"
var mongodb_collection = "producers"


// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/ping", pingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/producer", producerHandler(formatter)).Methods("GET")
	mx.HandleFunc("/producer/{pid}", producerInsertHandler(formatter)).Methods("POST")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Producer microservice is alive!"})
	}
}

// API Gumball Machine Handler
func producerHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//params := mux.Vars(req)
		//var ProducerId string = params["id"]

		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{"ProducerID": 1}).All(&result)
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Update Gumball Inventory
func producerInsertHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
        params := mux.Vars(req)
		var ProdId string = params["pid"]

		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		
		c := session.DB(mongodb_database).C(mongodb_collection)	
		var result bson.M
		err = c.Insert(bson.M{"ProducerId": ProdId})
        if err != nil {
                log.Fatal(err)
        }
		formatter.JSON(w, http.StatusOK, result)
	}
}
