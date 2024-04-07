package main

import (
	"context"
	"demo/gen-go/example" // 导入生成的 Thrift 代码
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"net"
)

func main() {
	// 设置服务地址和端口
	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9090"))
	if err != nil {
		log.Fatalln("Error creating socket:", err)
	}

	// 设置传输协议
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	// 设置传输方式
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	// 包装传输
	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		log.Fatalln("Error getting transport:", err)
	}

	// 打开传输
	if err := useTransport.Open(); err != nil {
		log.Fatalln("Error opening socket to 127.0.0.1:9090", " ", err)
	}
	defer useTransport.Close()

	// 创建客户端
	client := example.NewHelloWorldClientFactory(useTransport, protocolFactory)

	// 调用服务
	response, err := client.SayHello(context.Background(), "World")
	if err != nil {
		log.Fatalln("Error calling SayHello:", err)
	}
	log.Println("Response:", response)
}
