package shortener

import (
	"github.com/asaskevich/govalidator"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strings"
)

type JsonResponse map[string]interface{}

func GetApi() *martini.ClassicMartini {
	var storage Storage
	storage = NewRedisApi("localhost:6379", "") // add config

	api := martini.Classic()
	api.Use(render.Renderer(render.Options{
		Directory:  "static/views",
		Extensions: []string{".html"},
		Delims:     render.Delims{"{[{", "}]}"},
		IndentJSON: true,
	}))
	api.MapTo(storage, (*Storage)(nil))

	api.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	api.Post("/", func(r render.Render, storage Storage, req *http.Request) {
		if url := req.FormValue("url"); url != "" {
			if govalidator.IsURL(url) {
				if !strings.HasPrefix(url, "http") {
					url = "http://" + url
				}
				if hash, err := storage.CreateRecord(url); err != nil {
					r.JSON(http.StatusInternalServerError, JsonResponse{"error": err.Error()})
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
