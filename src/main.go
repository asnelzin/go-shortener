package main

import (
	"github.com/asnelzin/go-shortener/src/shortener"
	// "fmt"
	// "github.com/asaskevich/govalidator"
)

func main() {
	// fmt.Println(govalidator.IsURL(`www.google.com`))
	api := shortener.GetApi()
	api.Run()
}
