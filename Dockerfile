FROM golang:1.22.1-alpine3.19 AS build-stage
ADD . /app/
WORKDIR /app/
RUN go build -o bin/main cmd/main.go

FROM build-stage AS run-stage
WORKDIR /app/
EXPOSE 8080
ENTRYPOINT ["./bin/main"]