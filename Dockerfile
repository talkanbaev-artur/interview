#use small build environment
FROM golang:alpine as build 

WORKDIR /app

#get dependencies
COPY go.mod ./
RUN go mod download

#copy and build the project
COPY *.go ./
RUN go build -o /main

#Use deploy env
FROM alpine:latest

WORKDIR /

COPY --from=build /main /main

#optionally use --expose for dynamic config
EXPOSE 8000 

ENTRYPOINT [ "/main" ]