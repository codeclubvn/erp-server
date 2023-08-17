FROM golang:1.20-alpine as builder
WORKDIR /usr/src/app
COPY . .
ENV GOENV=production CGO_ENABLED=0 GOOS=linux
RUN go build -o erp-server .

FROM alpine:3.18 as production
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/erp-server .
CMD ["./erp-server"]
