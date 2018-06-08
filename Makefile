SHELL=/bin/bash

.PHONY: gen-go

gofiles = main.go client.go common.go web.go
thriftFiles = user

#下面这条命令可以获取到所有 thrift 文件的名字，但是目前还无法让bash执行结果输出到 Makefile 的数组中
#for file in $(ls ThriftFiles/*.thrift); do basename $file .thrift; done

gen-go:
	echo "Remove the vendor/user"
	rm -rf vendor/user

	thrift -r -gen go ThriftFiles/user.thrift

	echo "Move gen-go/user to the ./vendor/"
	mv gen-go/user ./vendor/

	rmdir gen-go

gen-py:
	rm -rvf UserService/user
	thrift -r -gen py ThriftFiles/user.thrift
	mv -v gen-py/user UserService
	rm -rvf gen-py

bin/rpc: $(gofiles) ThriftFiles
	govendor build -o bin/rpc $(gofiles)
