FROM golang:alpine 
MAINTAINER mi ma <mimamo@momima.de>
ADD /src /app/
WORKDIR /app
RUN go build -o server .
EXPOSE 80
ENTRYPOINT ["/app/server"]



