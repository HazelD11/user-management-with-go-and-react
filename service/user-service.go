package service

import (
	"github.com/AustinNick/coba-golang/model/domain"
	"github.com/AustinNick/coba-golang/model/web"
	"github.com/AustinNick/coba-golang/repository"
	"github.com/mashingan/smapping"
)

type UserService interface {
	All() []domain.User
	Create(b web.UserCreateRequest) (domain.User, error)
	// FindById(id uint) (domain.User, error)
	// Update(b web.UserUpdateRequest) (domain.User, error)
	// FindByEmail(email string) domain.User
	// Delete(id uint) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) All() []domain.User {
	return s.userRepository.FindAll()
}

func (s *userService) Create(req web.UserCreateRequest) (domain.User, error) {
	user := domain.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&req))
	if err != nil {
		return user, err
	}
	return s.userRepository.Create(user), nil
}
