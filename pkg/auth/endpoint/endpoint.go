package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/manan1979/watermark-service/pkg/auth"
)

type AuthSet struct {
	GetUserAccessEndpoint endpoint.Endpoint
	AuthenticateEndpoint  endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
}

func NewAuthEndpointSet(svc auth.Service) AuthSet {
	return AuthSet{
		GetUserAccessEndpoint: MakeGetUserAccessEndpoint(svc),
		AuthenticateEndpoint:  MakeAuthenticateEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
	}
}

func MakeGetUserAccessEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUserAccessRequest)
		userAcess, err := svc.GetUserAcess(ctx, req.Username)
		if err != nil {
			return GetUserAccessResponse{UserAccess: userAcess, Err: err.Error()}, nil
		}
		return GetUserAccessResponse{UserAccess: userAcess, Err: ""}, nil
	}
}

func MakeAuthenticateEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthenticateRequest)
		authorized, err := svc.Authenticate(ctx, req.Username, req.Operation)
		if err != nil {
			return AuthenticateResponse{Authorized: authorized, Err: err.Error()}, nil
		}
		return AuthenticateResponse{Authorized: authorized, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc auth.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}
