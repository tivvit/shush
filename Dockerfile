FROM golang:1.13

WORKDIR /app
ADD go.mod .
ADD go.sum .
ADD main.go .
ADD shush shush/

RUN go build -o shush-server main.go

CMD ["./shush-server"]