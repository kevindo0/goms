package main

import (
    "flag"
    "net/http"
    "log"
    "github.com/golang/glog"
    "golang.org/x/net/context"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"
    "google.golang.org/grpc"
    gw "goms/ggrpc/proto"
)

var (
    echoEndpoint = flag.String("echo_endpoint", "localhost:9192", "endpoint of Gateway")
)

func run() error {
    ctx := context.Background()
    ctx, cancel := context.WithCancel(ctx)
    defer cancel()

    mux := runtime.NewServeMux()
    opts := []grpc.DialOption{grpc.WithInsecure()}
    err := gw.RegisterGatewayHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
    if err != nil {
        return err
    }

    log.Println("服务开启")
    return http.ListenAndServe(":8080", mux)
}

func main() {
    flag.Parse()
    defer glog.Flush()

    if err := run(); err != nil {
        glog.Fatal(err)
    }
}