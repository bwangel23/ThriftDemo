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

func handleClient(client *tutorial.CalculatorClient) error {
	var err error

	client.Ping(defaultCtx)
	fmt.Println("ping()")

	// 执行一次加法操作
	var num1, num2 int32 = 2, 4
	sum, err := client.Add(defaultCtx, num1, num2)
	if err != nil {
		return err
	}
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)

	var comment string
	var work tutorial.Work
	var value int32

	// 通过 Work 执行一次乘法操作
	comment = "一次乘法操作"
	work.Op = tutorial.Operation_MULTIPLY
	work.Num2 = 20
	work.Num1 = 13
	work.Comment = &comment

	value, err = client.Calculate(defaultCtx, 1, &work)
	if err != nil {
		return err
	}
	fmt.Printf("%v[%s] = %d\n", work, *work.Comment, value)

	// 通过 Work 执行一次减法操作，并覆盖之前存储的操作结果
	comment = "一次减法操作"
	work.Op = tutorial.Operation_SUBTRACT
	work.Num2 = 20
	work.Num1 = 13
	work.Comment = &comment

	value, err = client.Calculate(defaultCtx, 1, &work)
	if err != nil {
		return err
	}
	fmt.Printf("%v[%s] = %d\n", work, *work.Comment, value)

	// 通过 logid 获取之前存储的操作结果
	calEntry, err := client.GetStruct(defaultCtx, 1)
	if err != nil {
		return err
	}
	fmt.Printf("Get calculate entry %s\n", calEntry)

	// 执行一次 zip 操作
	err = client.Zip(defaultCtx)
	if err != nil {
		return err
	}

	// 通过 Work 执行一次除法操作，并捕获异常
	comment = "一次无效的除法"
	work.Op = tutorial.Operation_DIVIDE
	work.Num1 = 1
	work.Num2 = 0
	work.Comment = &comment

	value, err = client.Calculate(defaultCtx, 2, &work)
	if err != nil {
		return err
	}
	fmt.Printf("%v[%s] = %d\n", work, *work.Comment, value)

	return nil
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
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

	return handleClient(tutorial.NewCalculatorClient(thrift.NewTStandardClient(iprotocol, oprotocol)))
}
