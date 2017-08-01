package main

import (
	"github.com/penguinn/pgo-test/JSONRPC2/handler"
	"github.com/penguinn/pgo/jsonrpc2"
	"github.com/penguinn/pgo"
	"log"
)

func main() {
	var controllers []jsonrpc2.Controller
	p := new(handler.Echo)
	controllers = append(controllers, jsonrpc2.Controller{Name:"Echo.Add", F:p.Add, Params:handler.AddParams{}, Result:handler.AddResult{}})
	controllers = append(controllers, jsonrpc2.Controller{Name:"Echo.Mul", F:p.Mul, Params:handler.MulParams{}, Result:handler.MulResult{}})

	err := pgo.Init("rpc.toml", controllers)
	if err != nil{
		log.Fatal(err)
	}
	err = pgo.Run()
	if err != nil{
		log.Fatal(err)
	}
}
