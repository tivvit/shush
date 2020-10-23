GIT_REVISION := `git rev-parse --short HEAD`
GIT_TAG := ``

build:
	go build -ldflags "-X github.com/tivvit/shush/cmd.GitTag=${GIT_TAG} -X github.com/tivvit/shush/cmd.GitCommit=${GIT_REVISION}" -o shush-server .

test:
	go test -v ./...

api-gen:
	docker-compose run swagger-gen
	mv go shush-api