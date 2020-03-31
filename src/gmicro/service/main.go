package main

import (
        "log"
        "github.com/micro/go-micro"
        proto "goms/gmicro/proto"
        "golang.org/x/net/context"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
    rsp.Greeting = "Hello " + req.Name
    return nil
}

func main() {
    service := micro.NewService(
        micro.Name("greeter"),
        micro.Version("latest"),
    )

    service.Init()

    proto.RegisterGreeterHandler(service.Server(), new(Greeter))

    if err := service.Run(); err != nil {
        log.Fatal(err)
    }
}