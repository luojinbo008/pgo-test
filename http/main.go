package main

import (
	"github.com/penguinn/pgo"
	"log"
	"github.com/penguinn/pgo-test/http/handler"
)

func main()  {
	var controller handler.Controller
	err := pgo.Init("web.toml", &controller)
	if err != nil{
		log.Fatal(err)
	}
	err = pgo.Run()
	if err != nil{
		log.Fatal(err)
	}
}