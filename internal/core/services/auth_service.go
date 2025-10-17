package services

import (
	"context"
	"errors"
	"strings"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
	"belajar-go/pkg/jwtutil"

	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	users ports.UserRepository
	jwt   jwtutil.TokenManager
}

func NewAuthService(users ports.UserRepository, jwt jwtutil.TokenManager) ports.AuthService {
	return &authService{users: users, jwt: jwt}
}

func (s *authService) Signup(ctx context.Context, email, password string) (*domain.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" || len(password) < 6 {
		return nil, errors.New("invalid email or password too short")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u := &domain.User{Email: email, Password: string(hashed)}
	if err := s.users.Create(ctx, u); err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}

func (s *authService) Login(ctx context.Context, email, password string) (string, *domain.User, error) {
	u, err := s.users.FindByEmail(ctx, strings.ToLower(strings.TrimSpace(email)))
	if err != nil || u == nil {
		return "", nil, errors.New("invalid email or password")
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return "", nil, errors.New("invalid email or password")
	}
	token, err := s.jwt.NewAccessToken(u.ID)
	if err != nil {
		return "", nil, err
	}
	u.Password = ""
	return token, u, nil
}
