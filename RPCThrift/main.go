package main

import (
	"github.com/penguinn/pgo-test/RPCThrift/thrift/RPCThrift"
	"github.com/penguinn/pgo-test/RPCThrift/handler"
	"github.com/penguinn/pgo"
	"log"
)

func main() {
	//handl := handler.ThriftHandler{}
	processor := RPCThrift.NewRPCThriftProcessor(new(handler.ThriftHandler))
	err := pgo.Init("thrift.toml", processor)
	if err != nil{
		log.Fatal(err)
	}
	err = pgo.Run()
	if err != nil{
		log.Fatal(err)
	}
}
