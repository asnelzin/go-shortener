package shortener

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

type JsonResponse map[string]interface{}

func GetApi() *martini.ClassicMartini {
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
			hash := storage.CreateRecord(url)
			response := JsonResponse{
				"shortUrl": hash,
			}
			r.JSON(http.StatusOK, response)
		}
	})

	api.Get("/(?P<hash>[a-zA-Z0-9]+)", func(r render.Render, params martini.Params, storage Storage, req *http.Request) {
		url := storage.GetUrl(params["hash"])
		r.Redirect(url)
	})

	return api
}
