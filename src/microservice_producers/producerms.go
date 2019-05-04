package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"os"
)

// MongoDB Config
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
	mx.HandleFunc("/producers", producerHandler(formatter)).Methods("GET")
	mx.HandleFunc("/producer/{pid}", producerPostHandler(formatter)).Methods("POST")
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Producer microservice is alive!"})
	}
}

// API Producer List Handler
func producerHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var result []bson.M
		err = c.Find(bson.M{}).All(&result)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, result)
	}
}

// API Producer List Handler
func producerPostHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		params := mux.Vars(req)
		var pid string = params["pid"]
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		var data = bson.D{{"producer", pid}}
		err = c.Insert(data)
		if err != nil {
			log.Fatal(err)
		}
		formatter.JSON(w, http.StatusOK, struct{ Message string }{"Success"})
	}
}

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3001"
	}

	server := NewServer()
	server.Run(":" + port)
}

/*
	use cmpe281;
    db.producers.insert({producer: 45261});

*/

