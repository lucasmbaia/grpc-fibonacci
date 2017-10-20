FROM golang
MAINTAINER grpc-fibonnaci
RUN mkdir /app
ADD grpc-fibonnaci /app/
WORKDIR /app
EXPOSE 5000
CMD ["/app/grpc-fibonnaci"]
