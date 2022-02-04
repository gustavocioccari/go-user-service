package user

import (
	"time"

	"github.com/gustavocioccari/go-user-microservice/models"
	"github.com/gustavocioccari/go-user-microservice/repositories/mongodb/user"
	"github.com/gustavocioccari/go-user-microservice/service/kafka"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	userRepository user.UserRepository
	kafkaService   kafka.KafkaService
}

type UserService interface {
	Create(user *models.User) (*models.User, error)
}

func NewUserService(
	userRepository user.UserRepository,
	kafkaService kafka.KafkaService,
) UserService {
	return &service{
		userRepository: userRepository,
		kafkaService:   kafkaService,
	}
}

func (s *service) Create(user *models.User) (*models.User, error) {
	user.ID = uuid.NewV4().String()
	user.CreatedAt, user.UpdatedAt = time.Now(), time.Now()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.Password = string(hash)

	userData, err := s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	userCreated, err := s.userRepository.FindById(userData.InsertedID.(string))
	if err != nil {
		return nil, err
	}

	userCreated.Password = ""

	s.kafkaService.Producer("new-user", userCreated)

	return userCreated, nil
}
