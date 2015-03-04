package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type JsonResponse map[string]interface{}

type Storage interface{}

func main() {
	var storage Storage
	redisStorage := NewRedisApi("localhost", "6379") // add config
	storage = redisStorage

	api := martini.Classic()
	api.Use(render.Renderer(render.Options{
		IndentJSON: true,
	}))
	api.MapTo(storage, (*Storage)(nil))

	api.Post("/", func(r render.Render, storage Storage, req *http.Request) {
		if url := req.FormValue("url"); url != "" {
			hash := redisStorage.createRecord(url)
			response := JsonResponse{
				"shortUrl": hash,
			}
			r.JSON(http.StatusOK, response)
		}
	})

	api.Run()
}
