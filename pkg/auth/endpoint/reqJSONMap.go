package endpoint

import "github.com/manan1979/watermark-service/pkg/auth"

type GetUserAccessRequest struct {
	Username string `json:"username"`
}

type GetUserAccessResponse struct {
	UserAccess auth.UserAcess `json:"user_access"`
	Err        string         `json:"err"`
}

type AuthenticateRequest struct {
	Username  string `json:"username"`
	Operation string `json:"operation"`
}

type AuthenticateResponse struct {
	Authorized bool   `json:"authorized"`
	Err        string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
