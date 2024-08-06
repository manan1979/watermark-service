package auth

import (
	"context"
	"errors"
	"net/http"
)

type authService struct{}

func NewService() Service {
	return &authService{}
}

func (s *authService) GetUserAcess(ctx context.Context, username string) (UserAcess, error) {
	if username == "admin" {
		return UserAcess{
			Roles:      []string{"admin"},
			Privileges: []string{"read", "write", "delete"},
		}, nil
	}
	return UserAcess{}, errors.New("user not found")
}

func (s *authService) Authenticate(ctx context.Context, username, operation string) (bool, error) {
	userAccess, err := s.GetUserAcess(ctx, username)
	if err != nil {
		return false, err
	}
	for _, privilege := range userAccess.Privileges {
		if privilege == operation {
			return true, nil
		}
	}
	return false, nil
}

func (s *authService) ServiceStatus(ctx context.Context) (int, error) {
	return http.StatusOK, nil
}
