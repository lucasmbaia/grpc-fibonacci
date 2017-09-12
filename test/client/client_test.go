package main

import (
  "context"

  "log"
  "testing"
  "github.com/lucasmbaia/grpc-fibonacci/client"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
)

func TestClientCalcFibonacci(t *testing.T) {
  var (
    err     error
    value   fibonacci.Result
    conf    client.Config
    number  fibonacci.Number
    ctx	    = context.Background()
  )

  number = fibonacci.Number{Value: 10}

  if value, err = conf.CalcFibonacci(ctx, &number); err != nil {
    log.Fatalf("Error to calc fibonacci: ", err)
  }

  log.Println(value.Value)
}

