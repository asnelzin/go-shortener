package main

import (
	"github.com/asnelzin/go-shortener/shortener"
)

func main() {
	api := shortener.GetApi()
	api.Run()
}
