package main

import (
	"github.com/asnelzin/shortener/src/shortener"
)

func main() {
	api := shortener.GetApi()
	api.Run()
}
