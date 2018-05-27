package main

import (
	"log"
	"git.apache.org/thrift.git/lib/go/thrift"

	// Thrift gen code
	"tutorial"

	"context"
	"fmt"
)

var defaultCtx = context.Background()

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	tSocket, err := thrift.NewTSocket(addr)
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}

	transport, err := transportFactory.GetTransport(tSocket)
	if err != nil {
		return err
	}

	client := tutorial.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		return err
	}
	defer transport.Close()

	data := tutorial.Data{Text: "hello, world!"}
	d, err := client.DoFormat(defaultCtx, &data)
	if err != nil {
		return err
	} else {
		fmt.Println(d.Text)
	}

	return nil
}
