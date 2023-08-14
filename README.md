# ERP Server
# Giải thích cấu trúc file
## Util
Nơi thực hiện lưu hằng số

Nơi thực hiện lưu các hàm hỗ trợ

## Model
Lưu trữ struct

## Config
Khởi tạo app
Connect DB
Config ENV

## Route
Định hướng URL tới function đích

## Handler
Xử lý request | response
`
Cách nhận biết: Có ctx *gin.Context là handler
ctx *gin.Context là một kiểu dữ liệu đại diện cho context của một request
Nó chứa các phương thức và thuốc tính cho phpes thao tác với request và response dễ dàng hơn
`


## Middleware
Bước trung gian giữa route và handler
VD: Check xem user có đủ quyền dùng app đó không

## Service
Xử lý logic chính

## Repository
Tương tác với database

## DB
Lưu trữ các file migration

# Giải thích (lý do và hướng chạy) các gói đã tạo trong dự án
### github.com/gin-gonic/gin
`
Gin được xây dựng dựa trên httprouter, một router nhanh và nhẹ được viết bằng Go.
Gin cung cấp một cách đơn giản để tạo các ứng dụng web hiệu suất cao và có thể mở rộng.
`

### gorm.io/gorm
`
GORM là một ORM (Object Relational Mapping) cho Golang.
`

### golang.org/x/crypto/bcrypt
`
Package bcrypt cung cấp hàm băm bcrypt cho Go (cách băm này khác
với MD5 ở chỗ .
`