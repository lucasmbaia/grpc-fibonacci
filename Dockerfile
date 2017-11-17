FROM golang
MAINTAINER fibonacci
RUN mkdir /app
ADD fibonacci /app/
ADD pkcs8.key /app/
ADD cacert.pem /app/
ADD nuvem-intera.local.pem /app/
WORKDIR /app
EXPOSE 5000
CMD ["/app/fibonacci"]
