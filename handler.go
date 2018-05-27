package main

import (
	"context"
	"fmt"
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
	fmt.Printf("IN: [%s] -- OUT: [%s]\n", data.Text, rData.Text)

	return &rData, nil
}
