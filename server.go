package main

import (
	"fmt"
	"tutorial"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	serverTransport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	handler := NewCalculatorHandler()
	processor := tutorial.NewCalculatorProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)

	fmt.Println("Running at:", addr)
	return server.Serve()
}
