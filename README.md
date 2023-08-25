# Golang Clean Architecture

## Cấu trúc thư mục

```
├── api
│ ├── controllers // Để xử lý các request từ client và trả về response
│ │ ├── auth.go
│ │ ├── base.go
│ │ ├── health.go
│ │ ├── module.go // Để khởi tạo các controller
│ │ └── user.go
│ ├── middlewares // Middleware để xử lý các request trước khi đến controller
│ │ ├── cors.go
│ │ ├── error_handle.go
│ │ ├── json.go
│ │ ├── jwt.go
│ │ ├── logger.go
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
│ ├── base.go
│ ├── role.go
│ ├── routes.go
│ └── user.go
├── README.md
├── repository // Để xử lý các thao tác với database
│ ├── module.go
│ └── user.go
├── service // Để xử lý các logic của project
│ ├── auth.go
│ ├── jwt.go
│ ├── module.go
│ └── user.go
└── utils // Các module để hỗ trợ cho việc xử lý request
```

## Các công nghệ sử dụng

- [Golang](https://golang.org/)
- [Gin](https://https://gin-gonic.com/)
- [Gorm](https://gorm.io/)
- [JWT](https://jwt.io/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)
