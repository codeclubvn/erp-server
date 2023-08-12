package repo

import (
	"pet-project/model"
)

type IUserRepo interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	UpdateUser(user model.User) (model.User, error)
	DeleteUser(user model.User) (model.User, error)
	GetUserById(id int) (model.User, error)
	GetAllUser() ([]model.User, error)
	GetRoleUser(role string) ([]model.User, error)
}

func (repo *Repo) GetUserByEmail(email string) (model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var user model.User
	repo.db.Where("email = ?", email).First(&user)
	return user, nil
}

func (repo *Repo) CreateUser(user model.User) (model.User, error) {
	// Tạo bản ghi mới trong bảng "users" từ request
	repo.db.Create(&user)
	return user, nil
}

func (repo *Repo) UpdateUser(user model.User) (model.User, error) {
	// Cập nhật bản ghi trong bảng "users" từ request
	repo.db.Save(&user)
	return user, nil
}

func (repo *Repo) DeleteUser(user model.User) (model.User, error) {
	// Xóa bản ghi trong bảng "users" từ request
	repo.db.Delete(&user)
	return user, nil
}

func (repo *Repo) GetUserById(id int) (model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var user model.User
	repo.db.Where("id = ?", id).First(&user)
	return user, nil
}

func (repo *Repo) GetAllUser() ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Find(&users)
	return users, nil
}

func (repo *Repo) GetRoleUser(role string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Where("role = ?", role).Find(&users)
	return users, nil
}

func (repo *Repo) GetStatusUser(status string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Where("status = ?", status).Find(&users)
	return users, nil
}

func (repo *Repo) GetAddressUser(address string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Where("address = ?", address).Find(&users)
	return users, nil
}

func (repo *Repo) GetCreateIdUser(create_id int) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Where("create_id = ?", create_id).Find(&users)
	return users, nil
}

func (repo *Repo) GetUpdateIdUser(update_id int) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Where("update_id = ?", update_id).Find(&users)
	return users, nil
}

func (repo *Repo) GetDeleteIdUser(delete_id int) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Where("delete_id = ?", delete_id).Find(&users)
	return users, nil
}

func (repo *Repo) GetDateOfBirth(ngaysinh string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var users []model.User
	repo.db.Where("date_of_birth = ?", ngaysinh).Find(&users)
	return users, nil
}

func (repo *Repo) GetPhoneUser(phone string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var user []model.User
	repo.db.Where("phone = ?", phone).Find(&user)
	return user, nil
}

func (repo *Repo) GetEmailUser(email string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var user []model.User
	repo.db.Where("email = ?", email).Find(&user)
	return user, nil
}

func (repo *Repo) GetUsernameUser(username string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var user []model.User
	repo.db.Where("username = ?", username).Find(&user)
	return user, nil
}

func (repo *Repo) GetPasswordUser(password string) ([]model.User, error) {
	// Tìm kiếm bản ghi trong bảng "users" từ request
	var user []model.User
	repo.db.Where("password = ?", password).Find(&user)
	return user, nil
}
