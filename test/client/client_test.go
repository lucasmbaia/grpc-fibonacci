package main

import (
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
  )

  number = fibonacci.Number{Value: 10}

  if value, err = conf.CalcFibonacci(&number); err != nil {
    log.Fatalf("Error to calc fibonacci: ", err)
  }

  log.Println(value.Value)
}

