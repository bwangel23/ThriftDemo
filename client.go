package main

import (
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"

	// Thrift gen code
	"user"

	"context"
	"time"
)

var defaultCtx = context.Background()

func handleClient(client *user.UserActivityClient) error {
	var err error

	now := time.Now()
	userid := int64(1)
	ev := user.ActivityEvent{
		Activity:  user.Activity_ONLINE,
		Timestamp: now.Unix(),
		Userid:    userid,
	}
	err = client.Online(defaultCtx, userid, &ev)
	log.Println("YES")
	if err != nil {
		return err
	}

	return err
}

func runClient(addr string) error {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tSocket, err := thrift.NewTSocket(addr)
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}

	transport, err := transportFactory.GetTransport(tSocket)
	if err != nil {
		return err
	}
	defer transport.Close()

	if err := transport.Open(); err != nil {
		return err
	}

	iprotocol := protocolFactory.GetProtocol(transport)
	oprotocol := protocolFactory.GetProtocol(transport)

	return handleClient(user.NewUserActivityClient(thrift.NewTStandardClient(iprotocol, oprotocol)))
}
