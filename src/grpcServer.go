package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"io/ioutil"
	"log"
	"net"
	. "rpc-gateway/src/api/service"
)

const (
	serverAddr = "127.0.0.1:8888" // Address gRPC服务地址
)

// 1、声明一个gRPCServer，里面是未实现的字段
type rpcServiceServer struct {
	UnimplementedGrpcGateWayServiceServer
}

// 加入服务端认证证书
func GetServeCreds() credentials.TransportCredentials {
	// TLS认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err := tls.LoadX509KeyPair("keys/server.pem", "keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to find server credentials %v", err)
	}

	certPool := x509.NewCertPool() // 初始化一个CertPool
	ca, err := ioutil.ReadFile("keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to find root credentials %v", err)
	}

	certPool.AppendCertsFromPEM(ca)          // 解析传入的证书，解析成功会将其加到池子中
	creds := credentials.NewTLS(&tls.Config{ // 构建基于TLS的TransportCredentials选项
		Certificates: []tls.Certificate{cert},        // 服务端证书链，可以有多个
		ClientAuth:   tls.RequireAndVerifyClientCert, // 要求必须验证客户端证书
		ClientCAs:    certPool,                       // 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
	})

	return creds
}

// 2、必须要实现在 ProdGrpcGateWay.proto 里声明的远程调用接口，否则客户端会报：
// rpc error: code = Unimplemented desc = method GetGrpcGateWayStock not implemente
func (s *rpcServiceServer) GetGrpcGateWayStock(ctx context.Context, in *GrpcGateWayRequest) (*GrpcGateWayResponse, error) {
	return &GrpcGateWayResponse{GoodsStock: in.GetGoodsId()}, nil
}

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(GetServeCreds()))         // 1、创建带ca证书验证的服务端
	RegisterGrpcGateWayServiceServer(rpcServer, &rpcServiceServer{}) // 2、注册服务
	listen, err := net.Listen("tcp", serverAddr)                     // 3、设置传输协议和监听地址
	if err != nil {
		log.Fatal("服务监听端口失败", err)
	}

	log.Println("rpcServer listen on " + serverAddr + " with TLS")
	rpcServer.Serve(listen) // 4、启动服务
}
