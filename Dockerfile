FROM golang
MAINTAINER fibonacci
RUN mkdir /app
ADD fibonacci /app/
WORKDIR /app
EXPOSE 5000
CMD ["/app/fibonacci"]
