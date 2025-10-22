package services

import (
	"context"
	"errors"
	"log"
	"regexp"
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

func (s *authService) Signup(ctx context.Context, email, username, password string) (*domain.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	username = strings.TrimSpace(strings.ToLower(username))

	if email == "" || username == "" || len(password) < 6 {
		return nil, errors.New("invalid signup payload")
	}
	if !regexp.MustCompile(`^[a-z0-9_\.]{3,20}$`).MatchString(username) {
		return nil, errors.New("invalid username (use a-z, 0-9, underscore, dot; 3-20 chars)")
	}

	if exist, _ := s.users.FindByUsername(ctx, username); exist != nil {
		return nil, errors.New("username already taken")
	}
	if exist, _ := s.users.FindByEmail(ctx, email); exist != nil {
		return nil, errors.New("email already used")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	u := &domain.User{Email: email, Username: username, Password: string(hashed), Role: domain.RoleMember}
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
	log.Printf("[DEBUG] user after repo: email=%s role=%q", u.Email, u.Role)
	u.Password = ""
	return token, u, nil
}
