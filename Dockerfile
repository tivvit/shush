FROM golang:1.13 as build
WORKDIR /app
ADD go.mod .
ADD go.sum .
ADD main.go .
ADD shush shush/
ADD cmd cmd/
ADD shush-api shush-api/
ADD server server/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o shush-server -ldflags "-X github.com/tivvit/shush/cmd.GitTag=$(git describe --tags 2>1 || echo "unknown") -X github.com/tivvit/shush/cmd.GitCommit=$(git rev-parse --short HEAD)" main.go

FROM alpine
WORKDIR /app
COPY --from=build /app/shush-server .
CMD ["./shush-server"]