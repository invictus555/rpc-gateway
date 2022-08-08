package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"log"
	"net/http"
	. "rpc-gateway/src/api/service"
)

const (
	serverAddr = "127.0.0.1:8888" // ServerAddress gRPC服务地址
	clientAddr = "127.0.0.1:8080" // ClientAddress 是浏览器等发送http请求时的地址
)

// 加入客户端认证证书
func GetClientCreds() credentials.TransportCredentials {
	// TLS连接
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err := tls.LoadX509KeyPair("keys/client.pem", "keys/client.key")
	if err != nil {
		grpclog.Fatalf("Failed to find client credentials %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("keys/ca.pem")
	if err != nil {
		grpclog.Fatalf("Failed to find root credentials %v", err)
	}

	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert}, // 客户端证书
		ServerName:   "ximu.info",             // 注意这里的参数为配置文件中所允许的ServerName，也就是其中配置的DNS...
		RootCAs:      certPool,
	})

	return creds
}

func main() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	mux := runtime.NewServeMux()                                                    // 1、创建路由
	opt := []grpc.DialOption{grpc.WithTransportCredentials(GetClientCreds())}       // 2、加入认证证书
	err := RegisterGrpcGateWayServiceHandlerFromEndpoint(ctx, mux, serverAddr, opt) // 3、注册请求服务端的方法
	if err != nil {
		log.Fatalf("cannot start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(clientAddr, mux) // 4、启动并监听http请求
	if err != nil {
		log.Fatalf("cannot listen and server: %v", err)
	}

	log.Println("httpServer listen on " + clientAddr + " with TLS")
}
