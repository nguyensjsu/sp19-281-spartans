package server

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"microservice/functions"
	"microservice/lib"
	"net/http"
)

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
	mx.HandleFunc("/temperature/{pid}/{startDate}/{endDate}", temperatureGetHandler(formatter)).Methods("GET")
}

// API Ping Handler
func pingHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"API alive"})
	}
}

// Get Temperature data
func temperatureGetHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		params := mux.Vars(req)
		var startDate string = params["startDate"]
		var endDate string = params["endDate"]
		var pid string = params["pid"]
		fmt.Println("startDate", startDate)
		fmt.Println("endDate", endDate)
		if startDate != "" && endDate != "" {
			fmt.Println(functions.GetValue1(startDate, endDate))
			formatter.JSON(w, http.StatusOK, lib.GetData(pid,startDate, endDate))
		} else {
			formatter.JSON(w, http.StatusOK, struct{ Message string }{"start and end dates are must"})
		}

		//if startDate != "" && endDate != "" {
		//	fmt.Println(functions.GetValue1(startDate, endDate))
		//	formatter.JSON(w, http.StatusOK, functions.GetValue1(startDate, endDate))
		//} else {
		//	formatter.JSON(w, http.StatusOK, struct{ Message string }{"start and end dates are must"})
		//}
	}
}
