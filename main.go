package main

import (
	"flag"
	"fmt"
	"os"
	"log"
	"git.apache.org/thrift.git/lib/go/thrift"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of RPCDemo:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	addr := flag.String("addr", "0.0.0.0:8080", "Address to listen to")
	server := flag.Bool("server", false, "Running Server")
	flag.Parse()


	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	if *server {
		// pass
	} else {
		if err := runClient(transportFactory, protocolFactory, *addr); err != nil {
			log.Fatalln("error running client", err)
		}
	}
}
