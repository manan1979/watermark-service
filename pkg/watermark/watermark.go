package watermark

import (
	"context"
	"net/http"
	"os"

	"github.com/manan1979/watermark-service/internal"
	"gorm.io/gorm"

	"github.com/lithammer/shortuuid/v3"

	"github.com/go-kit/kit/log"
)

type watermarkService struct {
	db *gorm.DB
}

type PaginationResponse struct {
	Documents   []internal.Document `json:"documents"`
	Total       int64               `json:"total"`
	CurrentPage int                 `json:"current_page"`
	TotalPages  int                 `json:"total_pages"`
}

func NewService(db *gorm.DB) Service { return &watermarkService{db: db} }

func (w *watermarkService) Get(ctx context.Context, page, pageSize int, filters ...internal.Filter) (PaginationResponse, error) {

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	var documents []internal.Document
	offset := (page - 1) * pageSize

	var total int64
	if err := w.db.Model(&internal.Document{}).Where(filters).Count(&total).Error; err != nil {
		return PaginationResponse{}, err
	}

	if err := w.db.Where(filters).Offset(offset).Limit(pageSize).Find(&documents).Error; err != nil {
		return PaginationResponse{}, err
	}

	totalPages := int((total + int64(pageSize) - 1) / int64(pageSize))
	return PaginationResponse{
		Documents:   documents,
		Total:       total,
		CurrentPage: page,
		TotalPages:  totalPages,
	}, nil

}

func (w *watermarkService) Status(_ context.Context, ticketID string) (internal.Status, error) {

	return internal.InProgress, nil
}

func (w *watermarkService) Watermark(_ context.Context, ticketID, mark string) (int, error) {

	return http.StatusOK, nil
}

func (w *watermarkService) AddDocument(_ context.Context, doc *internal.Document) (*internal.Document, error) {

	newTicketID := shortuuid.New()
	doc.TicketID = newTicketID

	if err := w.db.Create(&doc).Error; err != nil {
		return nil, err
	}
	return &internal.Document{
		TicketID:  doc.TicketID,
		Content:   doc.Content,
		Title:     doc.Title,
		Author:    doc.Author,
		Topic:     doc.Topic,
		Watermark: doc.Watermark,
	}, nil
}

func (w *watermarkService) ServiceStatus(_ context.Context) (int, error) {
	logger.Log("checking the service health...")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
