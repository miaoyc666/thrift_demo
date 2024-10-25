package main

import (
	"context"
	"demo/gen-go/test/com/hello"
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
)

// HelloWorldHandler 实现 HelloWorld 接口
type HelloWorldHandler struct{}

func (h *HelloWorldHandler) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello " + name, nil
}

func main() {
	// 设置监听端口
	transport, err := thrift.NewTServerSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		log.Fatalln("Error:", err)
	}

	// 设置传输协议，这里使用的是二进制协议
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	// 设置传输方式，这里使用的是缓冲传输
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	// 实例化服务处理器
	handler := &HelloWorldHandler{}
	processor := hello.NewHelloWorldProcessor(handler)

	// 创建服务端
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)

	log.Println("Starting the simple server... on ", "9090")
	err = server.Serve()
	if err != nil {
		log.Fatalln("Error:", err)
	}
}
