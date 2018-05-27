package main

import (
	"context"
	"log"
	"shared"
	"strconv"
	"strings"
	"tutorial"
)

type FormatDataImpl struct{}

func (fdi *FormatDataImpl) DoFormat(ctx context.Context, data *tutorial.Data) (r *tutorial.Data, err error) {
	// 这里 rData 因为在 DoFormat 返回时仍然有指针引用它，所以它不会被 GC 回收
	// 用 Go 语言的术语说就是 `rData 从函数 DoFormat 中逃逸了`
	// Go 语言会自动选择在栈上还是在堆上为局部变量分配空间，这个并不是由 var 或者 new 等声明方式决定的
	var rData tutorial.Data
	rData.Text = strings.ToUpper(data.Text)
	log.Printf("IN: [%s] -- OUT: [%s]\n", data.Text, rData.Text)

	return &rData, nil
}

type CalculatorHandler struct {
	log map[int32]*shared.SharedStruct
}

func NewCalculatorHandler() *CalculatorHandler {
	return &CalculatorHandler{
		log: make(map[int32]*shared.SharedStruct),
	}
}

func (self *CalculatorHandler) Ping(ctx context.Context) (err error) {
	log.Printf("ping()\n")
	return nil
}

func (self *CalculatorHandler) Calculate(ctx context.Context, logid int32, work *tutorial.Work) (val int32, err error) {
	// 打印调用记录
	log.Printf("Calculate logid: %d, work: %v[%s]\n", logid, work, *work.Comment)

	// 根据 work 的类型执行对应的操作
	//    如果是除0的话返回相应的异常
	switch work.Op {
	case tutorial.Operation_ADD:
		val = work.Num1 + work.Num2
	case tutorial.Operation_SUBTRACT:
		val = work.Num1 - work.Num2
	case tutorial.Operation_MULTIPLY:
		val = work.Num1 * work.Num2
	case tutorial.Operation_DIVIDE:
		if work.Num2 == 0 {
			ouch := &tutorial.InvalidOperation{
				Whatop: work.Op,
				Why:    "Cannot divide 0",
			}
			err = ouch
			return
		}
		val = work.Num1 / work.Num2
	default:
		ouch := tutorial.NewInvalidOperation()
		ouch.Whatop = work.Op
		ouch.Why = "Invalid Operation"
		err = ouch

		return
	}

	// 将这次执行的结果通过 logID 存储起来
	entry := shared.NewSharedStruct()
	entry.Key = logid
	entry.Value = strconv.Itoa(int(val))

	// 如果结果已经存在了，打印覆盖信息
	oldValue, exists := self.log[logid]
	if exists {
		log.Printf("Replace the %v with the %v\n", oldValue, entry)
	} else {
		log.Printf("Add entry %v for key %d\n", entry, logid)
	}

	self.log[logid] = entry

	// 返回结果和错误记录
	return val, err
}

func (self *CalculatorHandler) Add(ctx context.Context, num1 int32, num2 int32) (int32, error) {
	sum := num1 + num2
	log.Printf("%d + %d = %d\n", num1, num2, sum)
	return sum, nil
}

func (self *CalculatorHandler) GetStruct(ctx context.Context, key int32) (*shared.SharedStruct, error) {
	log.Printf("getStruct(%d)\n", key)

	v, exists := self.log[key]
	if !exists {
		log.Printf("%d not exists\n", key)
	}

	return v, nil
}

func (self *CalculatorHandler) Zip(ctx context.Context) (err error) {
	log.Println("zip()")
	return nil
}
