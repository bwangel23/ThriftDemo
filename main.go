package main

import (
	"flag"
	"fmt"
	"os"
	"log"
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of RPCDemo:\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n")
}

func main() {
	flag.Usage = Usage
	addr := flag.String("addr", "0.0.0.0:8080", "RPC Service Address")
	webAddr := flag.String("web-addr", "0.0.0.0:8000", "Web Socket Address to listen")
	flag.Parse()

	fmt.Println(*addr, *webAddr)

	if err := ServeWeb(*webAddr); err != nil {
		log.Fatalln("Serve Web", err)
	}
}
