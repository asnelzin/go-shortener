package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	redisConn *redis.Conn
)

func urlHash(conn redis.Conn) string {
	if key, err := conn.Do("INCR", "url.pointer"); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%x", key)
	}
}

func createRecord(conn redis.Conn, url string) string {
	hash := urlHash(conn)
	if _, err := conn.Do("SET", hash, url); err != nil {
		panic(err)
	}
	return hash
}

func getUrl(conn redis.Conn, hash string) string {
	if url, err := conn.Do("GET", hash); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%s", url)
	}
}

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
	// conn, err := redis.Dial("tcp", ":6379")
	// if err != nil {
	// 	panic(err)
	// }
	// defer conn.Close()

	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET", "POST")
	// router.HandleFunc("path", url)

	http.Handle("/", router)
	http.ListenAndServe(":3000", nil)
}
