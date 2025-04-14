FROM golang:1.24.2-alpine AS build

ENV GOPATH=/

WORKDIR /src

COPY ./ ./

RUN go mod download 

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o pvz_managment ./cmd/main.go

FROM alpine:latest

WORKDIR /root/configs

COPY --from=build /src/configs . 

WORKDIR /root/app

COPY --from=build /src/pvz_managment . 

CMD ["sh", "-c", "./pvz_managment -mode=${ENV}"]