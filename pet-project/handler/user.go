package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// gói thư viện dùng để mã hóa, giải mã mật khẩu
	"golang.org/x/crypto/bcrypt"

	"pet-project/config"
	"pet-project/model"
	"pet-project/service"
)

type User struct {
	service service.IUser
	user    model.User
}

type IUser interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	CheckPassword(password string) error
	HashPassword(password string) error
}

func NewUser(service service.IUser) *User {
	return &User{service: service}
}

func (h *User) Login(ctx *gin.Context) {
	// Lấy thông tin từ request
	userRequest := model.UserRequest{}
	if err := ctx.ShouldBindJSON(&userRequest); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Gọi service để xử lý logic
	userResponse, err := h.service.Login(ctx, userRequest)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, userResponse)
}

func (h *User) Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := h.HashPassword(user.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		// hàm ctx.Abort() được sử dụng để kết thúc việc xử lý
		// một yêu cầu HTTP hiện tại và ngừng bất kỳ xử lý tiếp theo nào trong chuỗi xử lý (middleware chain).
		ctx.Abort()
		return
	}

	record := config.DbDefault.Create(&user)
	if record.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": record.Error.Error(),
		})
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"usedID": user.ID, "email": user.Email, "address": user.Address, "role": user.Role, "status": user.Create_id, "createdAt": user.CreatedAt, "updatedAt": user.UpdatedAt, "deletedAt": user.DeletedAt})
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
