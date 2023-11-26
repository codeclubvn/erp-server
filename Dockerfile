# Stage 1: Install dependencies
FROM golang:1.21 as installer

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

# Stage 2: Build the application
FROM installer as builder

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./main.go

# Stage 3: Create a lightweight runtime image
FROM ubuntu:20.04 as runner

WORKDIR /app
COPY --from=builder /app/main /app/main
COPY --from=builder /app/config/config.yml /app/config/config.yml

RUN apt-get update \
  && DEBIAN_FRONTEND="noninteractive" apt-get -y install tzdata ca-certificates --no-install-recommends \
  && ln -fs /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
  && rm -fr /var/lib/apt/lists/*

EXPOSE $PORT

ENTRYPOINT ["/app/main"]
