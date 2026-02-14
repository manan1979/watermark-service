package watermark

import (
	"context"

	"github.com/manan1979/watermark-service/internal"
)

type Service interface {
	Get(ctx context.Context, page, pageSize int, filters ...internal.Filter) (PaginationResponse, error)
	Status(ctx context.Context, ticketID string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc *internal.Document) (*internal.Document, error)
	ServiceStatus(ctx context.Context) (int, error)
}
