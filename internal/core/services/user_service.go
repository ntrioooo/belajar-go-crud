package services

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"belajar-go/internal/core/domain"
	"belajar-go/internal/core/ports"
)

type userService struct {
	users ports.UserRepository
}

func NewUserService(users ports.UserRepository) ports.UserService {
	return &userService{users: users}
}

func (s *userService) GetMe(ctx context.Context, userID uint) (*domain.User, error) {
	if userID == 0 {
		return nil, errors.New("unauthorized")
	}
	u, err := s.users.FindByID(ctx, userID)
	if err != nil || u == nil {
		return nil, errors.New("user not found")
	}
	u.Password = ""
	return u, nil
}

func (s *userService) UpdateMe(ctx context.Context, userID uint, email, username *string) (*domain.User, error) {
	if userID == 0 {
		return nil, errors.New("unauthorized")
	}

	fields := map[string]any{}

	// normalisasi & validasi email (opsional: regex sederhana)
	if email != nil {
		e := strings.TrimSpace(strings.ToLower(*email))
		if e == "" {
			return nil, errors.New("email cannot be empty")
		}
		// cek unik
		exists, err := s.users.ExistsEmailExcept(ctx, e, userID)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("email already used")
		}
		fields["email"] = e
	}

	// normalisasi & validasi username
	if username != nil {
		u := strings.TrimSpace(strings.ToLower(*username))
		if u == "" {
			return nil, errors.New("username cannot be empty")
		}
		if !regexp.MustCompile(`^[a-z0-9_.]{3,20}$`).MatchString(u) {
			return nil, errors.New("invalid username (a-z, 0-9, _ . ; 3-20 chars)")
		}
		exists, err := s.users.ExistsUsernameExcept(ctx, u, userID)
		if err != nil {
			return nil, err
		}
		if exists {
			return nil, errors.New("username already taken")
		}
		fields["username"] = u
	}

	if len(fields) == 0 {
		// tidak ada yang diubah, kembalikan profil sekarang
		return s.GetMe(ctx, userID)
	}

	if err := s.users.UpdateByID(ctx, userID, fields); err != nil {
		return nil, err
	}

	return s.GetMe(ctx, userID)
}
