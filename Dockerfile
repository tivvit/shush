FROM golang:1.13 as build
WORKDIR /app
ADD go.mod .
ADD go.sum .
ADD main.go .
ADD shush shush/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shush-server main.go

FROM alpine
WORKDIR /app
COPY --from=build /app/shush-server .
CMD ["./shush-server"]