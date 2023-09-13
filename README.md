# Golang Clean Architecture

## Các công nghệ sử dụng

- [Golang](https://golang.org/)
- [Gin](https://https://gin-gonic.com/)
- [Gorm](https://gorm.io/)
- [JWT](https://jwt.io/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)

## Error handling

Use github.com/pkg/errors to wrap errors and add stack trace to errors

Error in repository layer must be wrapped by errors.Wrap(err,message)

Service layer don't need to wrap errors returned from repository layer, because it already wrapped. Service layer only need to wrap errors returned from it's own logic

## Config

Copy file config.example.yml to config.yml to config for local environment

## Cấu trúc thư mục

```
├── api
│ ├── controllers // Để xử lý các request từ client và trả về response
│ │ ├── health.go
│ │ ├── module.go // Để khởi tạo các controller
│ ├── middlewares // Middleware để xử lý các request trước khi đến controller
│ │ ├── middleware.go
│ │ └── module.go
│ └── response // Để trả về response cho client
│   └── response.go
├── api_errors
│ └── errors.go // Define các error code
├── bootstrap
│ └── bootstrap.go // Để khởi tạo các module
├── config
│ ├── config.dev.yml // Define các config cho môi trường dev
│ ├── config.go // Define các config
│ └── config.yml // Define các config cho môi trường local
├── constants // Define các constant cho toàn bộ project
│ ├── app.go
│ ├── platform.go
│ ├── role.go
│ └── token_type.go
├── Dockerfile
├── dto // Define các struct để validate request từ client
│ ├── auth
│ │ ├── login.go
│ │ └── register.go
│ └── user
├── go.mod
├── go.sum
├── infrastructure // Các module để kết nối với các service bên ngoài
│ ├── db.go
│ └── module.go
├── lib // Các module để hỗ trợ cho việc xử lý request
│ ├── logger.go
│ ├── module.go
│ └── server.go
├── main.go
├── models // Define các model để mapping với database
│ └── user.go
├── README.md
├── repository // Để xử lý các thao tác với database
│ ├── module.go
│ └── user.go
├── service // Để xử lý các logic của project
│ ├── module.go
│ └── user.go
└── utils // Các hàm hỗ trợ
```
