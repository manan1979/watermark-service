package endpoint

import "github.com/manan1979/watermark-service/internal"

type GetRequest struct {
	Filters  []internal.Filter `json:"filter,omitempty"`
	Page     int               `json:"page"`
	PageSize int               `json:"page_size"`
}

type GetResponse struct {
	Documents   []internal.Document `json:"documents"`
	Total       int64               `json:"total"`
	CurrentPage int                 `json:"current_page"`
	TotalPages  int                 `json:"total_pages"`
	Err         string              `json:"err,omitempty"`
}

type StatusRequest struct {
	TicketID string `json:"ticketID"`
}

type StatusResponse struct {
	Status internal.Status `json:"status"`
	Err    string          `json:"err,omitempty"`
}

type WatermarkRequest struct {
	TicketID string `json:"ticketID"`
	Mark     string `json:"mark"`
}

type WatermarkResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err,omitempty"`
}

type AddDocumentRequest struct {
	Document *internal.Document `json:"document"`
}

type AddDocumentResponse struct {
	TicketID  string `json:"ticketID"`
	Content   string `json:"context"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Topic     string `json:"topic"`
	Watermark string `json:"watermark,omitempty"`
	Err       string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
