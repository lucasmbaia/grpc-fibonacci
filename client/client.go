package client

import (
  "encoding/json"

  "context"
  contextGolang "golang.org/x/net/context"
  "google.golang.org/grpc"
  "github.com/lucasmbaia/grpc-base/config"
  "github.com/lucasmbaia/grpc-base/zipkin"
  "github.com/lucasmbaia/grpc-fibonacci/proto"
  "google.golang.org/grpc/credentials"
)

type Config struct {
  SSL bool
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
    collector zipkin.Collector
    span      zipkin.Span
  )

  if conn, err = c.connect(); err != nil {
    return value, err
  }
  defer conn.Close()

  if collector, span, err = c.openZipkin(ctx); err != nil {
    return value, err
  }
  defer collector.Conn.Close()

  span.Event([]string{"GRPC Client Receive"})
  span.Tag("Args Receive", c.convertArgs(n))

  cF = fibonacci.NewFibonacciServiceClient(conn)

  if r, err = cF.Calc(contextGolang.Background(), &fibonacci.Number{Value: n.Value}); err != nil {
    span.Span.Finish()
    return value, err
  }

  span.Event([]string{"GRPC Client Send"})
  span.Tag("Args Send", c.convertArgs(fibonacci.Result{Value: r.Value}))
  span.Span.Finish()

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

func (c Config) openZipkin(ctx context.Context) (zipkin.Collector, zipkin.Span, error) {
  var (
    collector zipkin.Collector
    span      zipkin.Span
    err	      error
    tags      = make(map[string]string)
  )

  if config.EnvConfig.Tracer {
    if collector, err = zipkin.NewCollector(config.EnvConfig.ZipkinURL, "0.0.0.0:0", config.EnvConfig.ServiceName, true); err != nil {
      return collector, span, err
    }

    tags = map[string]string{
      "Host":	  config.EnvConfig.ServiceIPs[0],
      "Hostname": config.EnvConfig.Hostname,
    }

    span = collector.OpenChildSpan(ctx, "Client", tags, nil)
  }

  return collector, span, nil
}

func (c Config) convertArgs(i interface{}) string {
  var (
    body  []byte
    err	  error
  )

  if body, err = json.Marshal(i); err != nil {
    return err.Error()
  }

  return string(body)
}
