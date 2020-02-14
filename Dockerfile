FROM golang:1.13

WORKDIR /app
ADD go.mod .
ADD main.go .
ADD shush .

RUN go build -o shush main.go

CMD ["./shush"]