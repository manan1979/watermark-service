package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/go-kit/kit/log"
	"github.com/manan1979/watermark-service/internal/util"
	"github.com/manan1979/watermark-service/pkg/database/endpoint"
)

func NewHTTPHandler(ep endpoint.Set) http.Handler {
	m := http.NewServeMux()

	m.Handle("/healthz", httptransport.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))
	m.Handle("/update", httptransport.NewServer(
		ep.UpdateEndpoint,
		decodeHTTPUpdateRequest,
		encodeResponse,
	))
	m.Handle("/add", httptransport.NewServer(
		ep.AddEndpoint,
		decodeHTTPAddRequest,
		encodeResponse,
	))
	m.Handle("/get", httptransport.NewServer(
		ep.GetEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
	))
	m.Handle("/remove", httptransport.NewServer(
		ep.RemoveEndpoint,
		decodeHTTPRemoveRequest,
		encodeResponse,
	))

	return m
}

func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.GetRequest
	if r.ContentLength == 0 {
		logger.Log("Get request with no body")
		return req, nil
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req endpoint.ServiceStatusRequest
	return req, nil
}

func decodeHTTPUpdateRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func decodeHTTPAddRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.AddRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPRemoveRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoint.RemoveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil

}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "applicatio/json; charset=utf-8")
	switch err {
	case util.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case util.ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
