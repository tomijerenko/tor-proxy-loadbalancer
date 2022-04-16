ARG ALPINE_VERSION=3.15
ARG GOLANG_VERSION=1.18
ARG COMPILER=golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION}

###################################################
##                    COMPILE                    ##
###################################################
FROM ${COMPILER} AS build
WORKDIR /app
COPY go.sum go.mod .
RUN go mod download
RUN go mod verify
COPY  . .
ARG CGO_ENABLED=0
RUN go build -ldflags "-w -s" -o loadbalancer cmd/loadbalancer/main.go

###################################################
##               CREATE CONTAINER                ##
###################################################
FROM alpine:${ALPINE_VERSION}

RUN apk update
RUN apk add --update docker openrc docker-compose
RUN rc-update add docker boot

WORKDIR /app
COPY --from=build /app/loadbalancer .
COPY --from=build /app/docker-compose.yaml .

EXPOSE 8080
ENTRYPOINT ["./loadbalancer"]
CMD ["8080"]
