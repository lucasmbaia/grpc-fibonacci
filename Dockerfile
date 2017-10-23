FROM golang
MAINTAINER grpc-fibonacci
RUN mkdir /app
ADD grpc-fibonacci /app/
WORKDIR /app
EXPOSE 5000
CMD ["/app/grpc-fibonacci"]
