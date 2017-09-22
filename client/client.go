package client

import (
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-base/zipkin"
  "github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
  "google.golang.org/grpc/credentials"
)

type Config struct {
  Collector zipkin.Collector
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

  if conn, err = c.connect(); err != nil {
    return value, err
  }
  defer conn.Close()

  cF = fibonacci.NewFibonacciServiceClient(conn)

  if r, err = cF.Calc(ctx, &fibonacci.Number{Value: n.Value}); err != nil {
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
      grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(c.Collector.Tracer)),
    }
  } else {
    opts = []grpc.DialOption{
      grpc.WithInsecure(),
      grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(c.Collector.Tracer)),
    }
  }

  return grpc.Dial(config.EnvLocal.LinkerdURL, opts...)
}
