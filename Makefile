all: clean openapi test build
openapi:
	curl https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar --output openapi-generator.jar
	java -jar openapi-generator.jar generate -g go-server -i myapi.yaml

build:
	CGO_ENABLED=0
	go get -d -v ./...
	go build -a -installsuffix cgo -o bin/main main.go

docker:
	docker build -t xavicampa/xavic-test . --no-cache

test: go
	@go test ./...

clean:
	rm -rf .openapi-generator
	rm -rf api
	rm -rf go
	rm -rf bin
	rm -f *.jar
mock:
	prism mock -p 8080 myapi.yaml

run: build
	./bin/main
