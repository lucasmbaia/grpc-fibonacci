package main

import (
  "log"
  "reflect"

  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-base/base"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
  "github.com/lucasmbaia/grpc-fibonacci/server"
)

func init() {
  config.LoadConfig()
}

func main() {
  var (
    configCMD base.ConfigCMD
    errChan   = make(chan error, 1)
  )

  go func() {
    configCMD = base.ConfigCMD {
      SSL:		true,
      ServiceServer:	reflect.Indirect(reflect.ValueOf(fibonacci.RegisterFibonacciServiceServer)),
      HandlerEndpoint:  reflect.Indirect(reflect.ValueOf(fibonacci.RegisterFibonacciServiceHandlerFromEndpoint)),
      ServerConfig:     server.NewFibonacciServer(),
    }

    errChan <- configCMD.Run()
  }()

  select {
  case e := <-errChan:
    log.Fatalf("Error grpc server: ", e)
  }
}
