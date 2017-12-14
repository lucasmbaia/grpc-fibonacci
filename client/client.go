package client

import (
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-base/base"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
)

type Config struct {
  base.Config
}

func init() {
  config.LoadConfig()
}

func (c Config) CalcFibonacci(ctx context.Context, n *fibonacci.Number) (fibonacci.Result, error) {
  var (
    conn      *grpc.ClientConn
    cF	      fibonacci.FibonacciServiceClient
    err	      error
    value     fibonacci.Result
    r	      *fibonacci.Result
  )

  if conn, err = c.ClientConnect(); err != nil {
    return value, err
  }
  defer conn.Close()

  cF = fibonacci.NewFibonacciServiceClient(conn)

  if r, err = cF.Calc(ctx, &fibonacci.Number{Value: n.Value}); err != nil {
    return value, err
  }

  return fibonacci.Result{Value: r.Value}, nil
}
