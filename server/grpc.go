package server

import (
  "golang.org/x/net/context"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
  empty	"github.com/golang/protobuf/ptypes/empty"
)

type FibonnaciServer struct {}

func NewFibonacciServer() FibonnaciServer {
  return FibonnaciServer{}
}

func (f FibonnaciServer) Calc(ctx context.Context, v *fibonacci.Number) (*fibonacci.Result, error) {
  var n = fib(v.Value)

  return &fibonacci.Result{Value: n}, nil
}

func (f FibonnaciServer) Health(ctx context.Context, emp *empty.Empty) (*empty.Empty, error) {
  select {
  case <-ctx.Done():
    return nil, ctx.Err()
  default:
    return new(empty.Empty), nil
  }
}

func fib(n int32) int32 {
  if n == 0 {
    return 0
  } else if n == 1 {
    return 1
  } else {
    return fib(n-1) + fib(n-2)
  }
}
