package main

import (
  "log"
  "testing"
  "github.com/lucasmbaia/grpc-fibonacci/client"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
  "github.com/lucasmbaia/grpc-base/zipkin"
  "github.com/lucasmbaia/grpc-base/base"
  "golang.org/x/net/context"
)

func TestClientCalcFibonacci(t *testing.T) {
  var (
    err	      error
    value     fibonacci.Result
    conf      client.Config
    number    fibonacci.Number
    ctx	      = context.Background()
    collector zipkin.Collector
  )

  if collector, err = newCollector() ; err != nil {
    log.Fatalf("Error to calc fibonacci: ", err)
  }

  number = fibonacci.Number{Value: 10}

  conf = client.Config {
    base.Config {
      Collector:	collector,
    },
  }

  if value, err = conf.CalcFibonacci(ctx, &number); err != nil {
    log.Fatalf("Error to calc fibonacci: ", err)
  }

  log.Println(value.Value)
}

func newCollector() (zipkin.Collector, error) {
  return zipkin.NewCollector(
    "http://172.16.95.113:9411/api/v1/spans",
    "192.168.75.128:9090",
    "fibonacci",
    true,
    false,
  )
}
