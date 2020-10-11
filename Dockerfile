FROM openjdk:alpine AS openapi
RUN wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/4.3.1/openapi-generator-cli-4.3.1.jar
WORKDIR /
COPY myapi.yaml .
RUN java -jar openapi-generator-cli-4.3.1.jar generate -g go-server -i myapi.yaml -o /out

FROM golang:1.15 AS build
RUN mkdir /go/src/myapi
WORKDIR /go/src/myapi
COPY --from=openapi /out/go ./go
#COPY go ./go
COPY myapi ./myapi
COPY main.go .
COPY go.mod .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime
COPY --from=build /go/src/myapi/openapi ./
EXPOSE 8080/tcp
ENTRYPOINT ["./openapi"]
