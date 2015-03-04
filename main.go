package main

import (
	"fmt"
	// "github.com/go-martini/martini"
	"github.com/gorilla/mux"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Write([]byte("Hello World"))
	case "POST":
		params := mux.Vars(r)
		rw.Write([]byte(r.FormValue("url")))
		rw.Write([]byte(params["url"]))
	default:
		panic("Not implemented")
	}
}

func main() {
	redisApi := NewRedisApi("localhost", "6379")
	fmt.Println(redisApi.createRecord("http://localhost"))

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET", "POST")

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}
