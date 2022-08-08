/**
* 测试与服务端通信是否正常的客户端
**/

package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	. "rpc-gateway/src/api/service"
)

const (
	Address = "127.0.0.1:8888" // Address gRPC服务地址
)

func main() {
	// 证书认证-双向认证
	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, _ := tls.LoadX509KeyPair("keys/client.pem", "keys/client.key")

	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile("keys/ca.pem")

	// 注意这里只能解析pem类型的根证书，所以需要的是ca.pem
	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	certPool.AppendCertsFromPEM(ca)

	// 构建基于 TLS 的 TransportCredentials 选项
	creds := credentials.NewTLS(&tls.Config{
		// 设置证书链，允许包含一个或多个
		Certificates: []tls.Certificate{cert},
		ServerName:   "ximu.info", //注意这里的参数为配置文件中所允许的ServerName，也就是其中配置的DNS...
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds)) // 1、建立连接
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	request := &GrpcGateWayRequest{
		GoodsId: 123,
	}

	query := NewGrpcGateWayServiceClient(conn)                           // 2. 调用 ProdGrpcGateWay_grpc.pb.go 中的 NewGrpcGateWayServiceClient 方法建立客户端
	res, err := query.GetGrpcGateWayStock(context.Background(), request) // 3、调用rpc方法
	if err != nil {
		log.Fatal("调用gRPC方法错误: ", err)
	}

	fmt.Println("grpc-gateway:调用gRPC方法成功,GoodsStock = ", res)

}
