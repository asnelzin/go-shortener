package main

import (
	"github.com/asnelzin/go-shortener/src/shortener"
)

func main() {
	api := shortener.GetApi()
	api.Run()
}
