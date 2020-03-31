package main

import (
    "fmt"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    pb "goms/ggrpc/proto" // 引入proto包
)

const (
    // Address gRPC服务地址
    Address = "127.0.0.1:9192"
)

func main() {
    // 连接
    conn, err := grpc.Dial(Address, grpc.WithInsecure())

    if err != nil {
        grpclog.Fatalln(err)
    }

    defer conn.Close()

    // 初始化客户端
    c := pb.NewGatewayClient(conn)

    // 调用方法
    reqBody := new(pb.StringMessage)
    reqBody.Value = "duc"
    r, err := c.Echo(context.Background(), reqBody)
    if err != nil {
        grpclog.Fatalln(err)
    }

    fmt.Println(r.Value)
}
