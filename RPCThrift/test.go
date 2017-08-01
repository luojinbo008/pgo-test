package main

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/penguinn/pgo-test/RPCThrift/handler"
	"github.com/penguinn/pgo-test/RPCThrift/thrift/RPCThrift"
)

func main() {
	var processor = thrift.NewTMultiplexedProcessor2()
	processor.RegisterProcessor("fitstProcessor", RPCThrift.NewRPCThriftProcessor(&handler.ThriftHandler{}))
}
