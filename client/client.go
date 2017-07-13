package client

import (
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
  "google.golang.org/grpc/credentials"
)

type Config struct {
  SSL bool
}

func init() {
  config.LoadConfig()
}

func (c Config) CalcFibonacci(n *fibonacci.Number) (fibonacci.Result, error) {
  var (
    conn  *grpc.ClientConn
    cF	  fibonacci.FibonacciServiceClient
    err	  error
    value fibonacci.Result
    r	  *fibonacci.Result
  )

  if conn, err = c.connect(); err != nil {
    return value, err
  }
  defer conn.Close()

  cF = fibonacci.NewFibonacciServiceClient(conn)

  if r, err = cF.Calc(context.Background(), &fibonacci.Number{Value: n.Value}); err != nil {
    return value, err
  }

  return fibonacci.Result{Value: r.Value}, nil
}

func (c Config) connect() (*grpc.ClientConn, error) {
  var (
    opts  []grpc.DialOption
    creds credentials.TransportCredentials
    err   error
  )

  if config.EnvConfig.GrpcSSL {
    if creds, err = credentials.NewClientTLSFromFile(config.EnvConfig.CAFile, config.EnvConfig.ServerNameAuthority); err != nil {
      return new(grpc.ClientConn), err
    }

    opts = []grpc.DialOption{
      grpc.WithTransportCredentials(creds),
    }
  } else {
    opts = []grpc.DialOption{
      grpc.WithInsecure(),
    }
  }

  return grpc.Dial(config.EnvLocal.LinkerdURL, opts...)
}

