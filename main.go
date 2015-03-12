package main

import (
	"shortener"
)

func main() {
	api := shortener.GetApi()
	api.Run()
}
