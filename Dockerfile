FROM golang:1.21-alpine as builder
WORKDIR /usr/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN GOENV=production CGO_ENABLED=0 GOOS=linux go build -o erp-server .

FROM alpine:3.18.3 as production
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/erp-server .
CMD ["./erp-server"]
