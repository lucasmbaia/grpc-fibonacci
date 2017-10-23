FROM golang
MAINTAINER fibonnaci
RUN mkdir /app
ADD fibonnaci /app/
WORKDIR /app
EXPOSE 5000
CMD ["/app/fibonnaci"]
