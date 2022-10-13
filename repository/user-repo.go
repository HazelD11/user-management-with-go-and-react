package repository

import (
	"github.com/AustinNick/coba-golang/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []domain.User
	Save(u domain.User) domain.User
	// FindById(id uint) (domain.User, error)
	// Update(b domain.User) (domain.User, error)
	// FindByEmail(email string) domain.User
	// Delete(id uint) error
}

type UserConnection struct {
	con *gorm.DB
}

func NewUserRepository(con *gorm.DB) UserRepository {
	return &UserConnection{con: con}
}

func (db *UserConnection) FindAll() []domain.User {
	var users []domain.User
	db.con.Find(&users)
	return users
}

func (db *UserConnection) Save(user domain.User) domain.User {
	db.con.Save(&user)
	return user
}
