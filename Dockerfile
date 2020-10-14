FROM openjdk:alpine AS openapi
RUN set -eux; \
	apk add --no-cache --virtual make
RUN set -eux; \
	apk add --no-cache --virtual curl
WORKDIR /
COPY Makefile .
COPY myapi.yaml .
RUN make openapi

FROM golang:1.15 AS build
RUN mkdir /go/src/myapi
WORKDIR /go/src/myapi
COPY --from=openapi /go go
COPY --from=openapi /api api
COPY Makefile .
COPY myapi ./myapi
COPY main.go .
COPY go.mod .
ENV CGO_ENABLED=0
RUN make build

FROM scratch AS runtime
COPY --from=build /go/src/myapi/bin/main ./
EXPOSE 8080/tcp
ENTRYPOINT ["./main"]
