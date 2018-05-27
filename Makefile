.PHONY: gen-go

gofiles = main.go server.go client.go handler.go

gen-go:
	rm -rf vendor/tutorial vendor/shared
	thrift -r -gen go ThriftFiles/tutorial.thrift
	mv gen-go/* ./vendor/
	rmdir gen-go

bin/rpc: $(gofiles) ThriftFiles
	govendor build -o bin/rpc $(gofiles)
