package auth

import "context"

type UserAcess struct {
	Roles      []string `json:"roles"`
	Privileges []string `json:"priviliges"`
}

type Service interface {
	GetUserAcess(ctx context.Context, username string) (UserAcess, error)
	Authenticate(ctx context.Context, username, operation string) (bool, error)
	ServiceStatus(ctx context.Context) (int, error)
}
