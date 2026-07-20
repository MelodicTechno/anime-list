package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"github.com/MelodicTechno/anime-list/internal/model"
	"github.com/MelodicTechno/anime-list/internal/repository"
)

var (
	ErrUsernameTaken      = errors.New("username already exists")
	ErrEmailTaken         = errors.New("email already exists")
	ErrInvalidCredential  = errors.New("invalid username or password")
	ErrWeakPassword       = errors.New("password must be at least 6 characters")
	ErrInvalidEmail       = errors.New("invalid email format")
	ErrInvalidRefreshToken = errors.New("invalid or expired refresh token")
)

type UserService struct {
	repo              *repository.UserRepository
	rdb               *redis.Client
	jwtSecret         []byte
	accessExpireHours int
	refreshExpireHours int
}

func NewUserService(repo *repository.UserRepository, rdb *redis.Client, jwtSecret string, accessHours, refreshHours int) *UserService {
	return &UserService{
		repo:              repo,
		rdb:               rdb,
		jwtSecret:         []byte(jwtSecret),
		accessExpireHours: accessHours,
		refreshExpireHours: refreshHours,
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

func (s *UserService) Login(account, password string) (string, string, *model.User, error) {
	account = strings.TrimSpace(account)
	if account == "" || password == "" {
		return "", "", nil, ErrInvalidCredential
	}

	user, err := s.repo.FindByUsernameOrEmail(account)
	if err != nil {
		return "", "", nil, err
	}
	if user == nil {
		return "", "", nil, ErrInvalidCredential
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", "", nil, ErrInvalidCredential
	}

	accessToken, err := s.generateAccessToken(user.ID)
	if err != nil {
		return "", "", nil, err
	}

	refreshToken, err := s.generateRefreshToken(user.ID)
	if err != nil {
		return "", "", nil, err
	}

	return accessToken, refreshToken, user, nil
}

func (s *UserService) RefreshToken(refreshToken string) (string, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return s.jwtSecret, nil
	})
	if err != nil {
		return "", ErrInvalidRefreshToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", ErrInvalidRefreshToken
	}

	tokenType, _ := claims["type"].(string)
	if tokenType != "refresh" {
		return "", ErrInvalidRefreshToken
	}

	userIDFloat, ok := claims["userId"].(float64)
	if !ok {
		return "", ErrInvalidRefreshToken
	}
	userID := int64(userIDFloat)

	redisKey := fmt.Sprintf("refresh:%d", userID)
	stored, err := s.rdb.Get(context.Background(), redisKey).Result()
	if errors.Is(err, redis.Nil) || stored != refreshToken {
		return "", ErrInvalidRefreshToken
	}
	if err != nil {
		return "", err
	}

	return s.generateAccessToken(userID)
}

func (s *UserService) generateAccessToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"type":   "access",
		"exp":    time.Now().Add(time.Duration(s.accessExpireHours) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.jwtSecret)
}

func (s *UserService) generateRefreshToken(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"userId": userID,
		"type":   "refresh",
		"exp":    time.Now().Add(time.Duration(s.refreshExpireHours) * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", err
	}

	redisKey := fmt.Sprintf("refresh:%d", userID)
	expire := time.Duration(s.refreshExpireHours) * time.Hour
	if err := s.rdb.Set(context.Background(), redisKey, signed, expire).Err(); err != nil {
		return "", err
	}

	return signed, nil
}
