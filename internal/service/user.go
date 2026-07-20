package service

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/MelodicTechno/anime-list/internal/model"
	"github.com/MelodicTechno/anime-list/internal/repository"
)

var (
	ErrUsernameTaken     = errors.New("username already exists")
	ErrEmailTaken        = errors.New("email already exists")
	ErrInvalidCredential = errors.New("invalid username or password")
	ErrWeakPassword      = errors.New("password must be at least 6 characters")
	ErrInvalidEmail      = errors.New("invalid email format")
)

type UserService struct {
	repo       *repository.UserRepository
	jwtSecret  []byte
	jwtExpire  time.Duration
}

func NewUserService(repo *repository.UserRepository, jwtSecret string, expireHours int) *UserService {
	return &UserService{
		repo:      repo,
		jwtSecret: []byte(jwtSecret),
		jwtExpire: time.Duration(expireHours) * time.Hour,
	}
}

func (s *UserService) Register(username, email, password string) (*model.User, error) {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)

	if len(username) == 0 {
		return nil, ErrInvalidCredential
	}
	if len(password) < 6 {
		return nil, ErrWeakPassword
	}
	if !strings.Contains(email, "@") {
		return nil, ErrInvalidEmail
	}

	existing, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrUsernameTaken
	}

	existing, err = s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, ErrEmailTaken
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hash),
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(account, password string) (string, *model.User, error) {
	account = strings.TrimSpace(account)
	if account == "" || password == "" {
		return "", nil, ErrInvalidCredential
	}

	user, err := s.repo.FindByUsernameOrEmail(account)
	if err != nil {
		return "", nil, err
	}
	if user == nil {
		return "", nil, ErrInvalidCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, ErrInvalidCredential
	}

	token, err := s.generateToken(user.ID)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

func (s *UserService) generateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"exp":    time.Now().Add(s.jwtExpire).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}
