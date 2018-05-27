package main

import (
	"fmt"
	"tutorial"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	handler := &FormatDataImpl{}
	processor := tutorial.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", addr)
	err = server.Serve()
	if err != nil {
		return err
	}

	return nil
}
