# syntax=docker/dockerfile:1
# Could use golang/alpine image
# FROM alpine:3.17 -- Decided golang instead because "RUN" commands create layers (larger size) and we use less this way
FROM golang:bullseye

# Set WORKDIR to prevent errors
WORKDIR /go/delivery

# Could clone from git instead
COPY . ./

RUN go build -o prober ./cmd/prober/main.go
ENTRYPOINT ["./prober"]

# CMD -config ./config/config.example.yaml -- https://docs.docker.com/develop/develop-images/dockerfile_best-practices/#cmd
CMD ["-config", "./config/config.example.yaml"]


# TO BUILD IMAGE
# docker build -t prober-docker . -f build/package/Dockerfile

# TO USE OTHER CONFIG FILES CAN BE MOUNTED PREVIOUSLY TO THE CONTAINER
# docker run -v /path/to/config.yaml:/path/in/container/to/config.yaml docker-prober --config /path/in/container/to/config.yaml