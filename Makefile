SHELL=/bin/bash

.PHONY: gen-go

gofiles = main.go server.go client.go handler.go common.go
thriftFiles = shared tutorial

#下面这条命令可以获取到所有 thrift 文件的名字，但是目前还无法让bash执行结果输出到 Makefile 的数组中
#for file in $(ls ThriftFiles/*.thrift); do basename $file .thrift; done

gen-go:
	@for thriftFile in $(thriftFiles); do\
		echo "Remove the vendor/$$thriftFile"; \
		rm -rf vendor/$$thriftFile; \
	done

	thrift -r -gen go ThriftFiles/tutorial.thrift

	@for thriftFile in $(thriftFiles); do\
		echo "Move gen-go/$$thriftFile/ to the ./vendor/"; \
		mv gen-go/$$thriftFile/ ./vendor/; \
	done
	rmdir gen-go

bin/rpc: $(gofiles) ThriftFiles
	govendor build -o bin/rpc $(gofiles)
