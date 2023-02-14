#Building
FROM golang:1.17-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
WORKDIR /app/cmd/prober
RUN go build -o prober .

#Running
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/cmd/prober/prober .
COPY --from=build /app/config/config.example.yaml /etc/prober/config.yaml
EXPOSE 2112
CMD ["./prober"]


