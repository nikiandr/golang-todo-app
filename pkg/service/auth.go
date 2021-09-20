package service

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/nikiandr/golang-todo-app"
	"github.com/nikiandr/golang-todo-app/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//JWT token validity in hours
const (
	jwtValidityTime = 24
	signingKey      = "d%jfoi2@e12u89$yurui#13fheh"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	password, err := s.generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = password
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from DB
	user, err := s.repo.GetUser(username)
	if err != nil {
		logrus.Errorf("Error getting user from DB: %s", err.Error())
		return "", err
	}

	logrus.Println(user.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logrus.Errorf("Wrong password: %s", err.Error())
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtValidityTime * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Errorf("Error hashing password: %s", err.Error())
		return "", err
	}
	return string(hash), nil
}
