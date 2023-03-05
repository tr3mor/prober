# syntax=docker/dockerfile:1
# Could use golang/alpine image
# FROM alpine:3.17 -- Decided golang instead because "RUN" commands create layers (larger size) and we use less this way
# If using alpine --> Should use multi-stage build https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#use-multi-stage-builds
FROM golang:1.20.1-alpine3.17 as builder

# Set WORKDIR to prevent errors
WORKDIR /go/delivery

# Could clone from git instead
# COPY go.mod go.sum cmd ./
COPY . .

RUN go build -o prober ./cmd/prober/main.go

FROM alpine:3.17.2
WORKDIR /prober
COPY --from=builder /go/delivery/config /go/delivery/prober ./
CMD ["./prober", "-config", "./config.yaml"]

# TO USE OTHER CONFIG FILES CAN BE MOUNTED PREVIOUSLY TO THE CONTAINER
# docker run -v /path/to/config.yaml:/path/in/container/to/config.yaml stasky745/prober ./prober --config /path/in/container/to/config.yaml