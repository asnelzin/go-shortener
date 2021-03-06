package shortener

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"os"
	"strings"
)

type JsonResponse map[string]interface{}

func GetApi() *martini.ClassicMartini {
	var storage Storage
	var redisHost, redisPassword string
	if redisHost = os.Getenv("REDIS_HOST"); redisHost == "" {
		redisHost = "localhost:6379"
	}
	if redisPassword = os.Getenv("REDIS_PASSWORD"); redisPassword == "" {
		redisPassword = ""
	}
	storage = NewRedisApi(redisHost, redisPassword) // add config

	api := martini.Classic()
	api.Use(render.Renderer(render.Options{
		Directory:  "static/views",
		Extensions: []string{".html"},
		Delims:     render.Delims{"{[{", "}]}"},
		IndentJSON: true,
	}))
	if os.Getenv("GO_DEBUG") == "true" {
		api.Use(martini.Static("."))
	}
	api.MapTo(storage, (*Storage)(nil))

	api.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	api.Post("/", func(r render.Render, storage Storage, req *http.Request) {
		if url := req.FormValue("url"); url != "" {
			if IsURL(url) {
				if !strings.HasPrefix(url, "http") {
					url = "http://" + url
				}
				if hash, err := storage.CreateRecord(url); err != nil {
					r.JSON(http.StatusInternalServerError, JsonResponse{"error": "Redis connection refused"})
				} else {
					r.JSON(http.StatusOK, JsonResponse{"shortUrl": hash})
				}
			} else {
				r.JSON(http.StatusBadRequest, JsonResponse{"error": "URL is not valid"})
			}
		}
	})

	api.Get("/(?P<hash>[a-zA-Z0-9]+)", func(r render.Render, params martini.Params, storage Storage, req *http.Request) {
		if url, err := storage.GetUrl(params["hash"]); err != nil {
			r.JSON(http.StatusNotFound, JsonResponse{"error": err.Error()})
		} else {
			r.Redirect(url)
		}
	})

	return api
}
