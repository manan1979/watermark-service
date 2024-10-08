package endpoint

import (
	"context"
	"errors"
	"os"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/manan1979/watermark-service/internal"
	"github.com/manan1979/watermark-service/pkg/watermark"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddDocumentEndpoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	WatermarkEndpoint     endpoint.Endpoint
}

type PaginationResponse struct {
	Documents   []internal.Document `json:"documents"`
	Total       int64               `json:"total"`
	CurrentPage int                 `json:"current_page"`
	TotalPages  int                 `json:"total_pages"`
}

func NewEndpointSet(svc watermark.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		AddDocumentEndpoint:   MakeAddDocumentEndpoint(svc),
		StatusEndpoint:        MakeStatusEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
		WatermarkEndpoint:     MakeWatermarkEndpoint(svc),
	}
}

func MakeGetEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRequest)
		response, err := svc.Get(ctx, req.Page, req.PageSize, req.Filters...)
		if err != nil {
			return GetResponse{
				Documents:   nil,
				Total:       0,
				CurrentPage: req.Page,
				TotalPages:  0,
				Err:         err.Error(),
			}, nil
		}
		return GetResponse{
			Documents:   response.Documents,
			Total:       response.Total,
			CurrentPage: req.Page,
			TotalPages:  response.TotalPages,
			Err:         "",
		}, nil
	}

}

func MakeStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*StatusRequest)
		status, err := svc.Status(ctx, req.TicketID)
		if err != nil {
			return StatusResponse{Status: status, Err: err.Error()}, nil
		}
		return StatusResponse{Status: status, Err: ""}, nil
	}
}

func MakeAddDocumentEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*AddDocumentRequest)
		ticketID, err := svc.AddDocument(ctx, req.Document)
		if err != nil {
			return AddDocumentResponse{
				TicketID:  ticketID.TicketID,
				Content:   ticketID.Content,
				Title:     ticketID.Author,
				Topic:     ticketID.Title,
				Author:    ticketID.Author,
				Watermark: ticketID.Watermark,
				Err:       err.Error()}, nil
		}
		return AddDocumentResponse{
			TicketID:  ticketID.TicketID,
			Content:   ticketID.Content,
			Title:     ticketID.Author,
			Topic:     ticketID.Title,
			Author:    ticketID.Author,
			Watermark: ticketID.Watermark,
		}, nil
	}
}

func MakeWatermarkEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(*WatermarkRequest)
		code, err := svc.Watermark(ctx, req.TicketID, req.Mark)
		if err != nil {
			return WatermarkResponse{Code: code, Err: err.Error()}, nil
		}
		return WatermarkResponse{Code: code, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code, err := svc.ServiceStatus(ctx)
		if err != nil {
			return ServiceStatusResponse{Code: code, Err: err.Error()}, nil
		}
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context, page, pageSize int, filters ...internal.Filter) (PaginationResponse, error) {

	resp, err := s.GetEndpoint(ctx, GetRequest{
		Filters:  filters,
		Page:     page,
		PageSize: pageSize,
	})
	if err != nil {
		return PaginationResponse{}, err
	}

	getResp := resp.(GetResponse)
	if getResp.Err != "" {
		return PaginationResponse{}, errors.New(getResp.Err)
	}

	return PaginationResponse{
		Documents:   getResp.Documents,
		Total:       getResp.Total,
		CurrentPage: getResp.CurrentPage,
		TotalPages:  getResp.TotalPages,
	}, nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {

	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}

	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

func (s *Set) AddDocument(ctx context.Context, doc *internal.Document) (string, error) {
	resp, err := s.AddDocumentEndpoint(ctx, AddDocumentRequest{Document: doc})
	if err != nil {
		return "", err
	}
	adResp := resp.(AddDocumentResponse)
	if adResp.Err != "" {
		return "", errors.New(adResp.Err)
	}
	return adResp.TicketID, nil

}

func (s *Set) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	resp, err := s.StatusEndpoint(ctx, StatusRequest{TicketID: ticketID})
	if err != nil {
		return internal.Failed, err
	}

	stsResp := resp.(StatusResponse)
	if stsResp.Err != "" {
		return internal.Failed, errors.New(stsResp.Err)
	}
	return stsResp.Status, nil
}

func (s *Set) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	resp, err := s.WatermarkEndpoint(ctx, WatermarkRequest{TicketID: ticketID, Mark: mark})
	wmResp := resp.(WatermarkResponse)
	if err != nil {
		return wmResp.Code, err
	}

	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}
	return wmResp.Code, nil

}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC())
}
