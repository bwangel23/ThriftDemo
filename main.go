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
	protocol := flag.String("P", "binary", "Specify the protocol (binary, compact, json, simplejson)")
	buffered := flag.Bool("buffered", false, "Use the buffered transport")
	framed := flag.Bool("framed", false, "Use framed transport")
	flag.Parse()


	var protocolFactory thrift.TProtocolFactory
	switch *protocol {
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified", protocol)
		Usage()
		os.Exit(1)
	}

	var transportFactory thrift.TTransportFactory
	if *buffered {
		transportFactory = thrift.NewTBufferedTransportFactory(8192)
	} else {
		transportFactory = thrift.NewTTransportFactory()
	}

	if *framed {
		transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	}

	if *server {
		if err := runServer(transportFactory, protocolFactory, *addr); err != nil {
			log.Fatalln("error running server", err)
		}
	} else {
		if err := runClient(transportFactory, protocolFactory, *addr); err != nil {
			log.Fatalln("error running client", err)
		}
	}
}
