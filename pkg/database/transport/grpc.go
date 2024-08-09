package transport

import (
	"context"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/manan1979/watermark-service/api/pb/db"
	"github.com/manan1979/watermark-service/internal"
	"github.com/manan1979/watermark-service/pkg/database/endpoint"
)

type grpcServer struct {
	db.DatabaseServer
	get           grpctransport.Handler
	update        grpctransport.Handler
	add           grpctransport.Handler
	remove        grpctransport.Handler
	serviceStatus grpctransport.Handler
}

func NewGRPCServer(ep endpoint.Set) db.DatabaseServer {

	return &grpcServer{
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
		add: grpctransport.NewServer(
			ep.AddEndpoint,
			decodeGRPCAddRequest,
			decodeGRPCAddResponse,
		),
		update: grpctransport.NewServer(
			ep.UpdateEndpoint,
			decodeGRPCUpdateRequest,
			decodeGRPCUpdateResponse,
		),
		remove: grpctransport.NewServer(
			ep.RemoveEndpoint,
			decodeGRPCRemoveRequest,
			decodeGRPCRemoveResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}

func (g *grpcServer) Get(ctx context.Context, r *db.GetRequest) (*db.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.GetReply), nil
}

func (g *grpcServer) Add(ctx context.Context, r *db.AddRequest) (*db.AddReply, error) {
	_, rep, err := g.add.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.AddReply), nil
}

func (g *grpcServer) Update(ctx context.Context, r *db.UpdateRequest) (*db.UpdateReply, error) {
	_, rep, err := g.update.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.UpdateReply), nil
}

func (g *grpcServer) Remove(ctx context.Context, r *db.RemoveRequest) (*db.RemoveReply, error) {
	_, rep, err := g.update.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.RemoveReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, r *db.ServiceStatusRequest) (*db.ServiceStatusReply, error) {
	_, rep, err := g.serviceStatus.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*db.ServiceStatusReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.GetRequest)
	var filters []internal.Filter
	for _, f := range req.Filters {
		filters = append(filters, internal.Filter{Key: f.Key, Value: f.Value})
	}
	return endpoint.GetRequest{Filters: filters}, nil
}

func decodeGRPCAddRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.AddRequest)
	doc := &internal.Document{
		Content:   req.Document.Content,
		Title:     req.Document.Title,
		Author:    req.Document.Author,
		Topic:     req.Document.Topic,
		Watermark: req.Document.Watermark,
	}
	return endpoint.AddRequest{Document: doc}, nil
}

func decodeGRPCUpdateRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.UpdateRequest)
	doc := &internal.Document{
		Content:   req.Document.Content,
		Title:     req.Document.Title,
		Author:    req.Document.Author,
		Topic:     req.Document.Topic,
		Watermark: req.Document.Watermark,
	}
	return endpoint.UpdateRequest{TicketID: req.TicketID, Document: doc}, nil
}

func decodeGRPCRemoveRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*db.RemoveRequest)
	return endpoint.RemoveRequest{TicketID: req.TicketID}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, _ interface{}) (interface{}, error) {
	return endpoint.ServiceStatusRequest{}, nil
}

func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.GetReply)
	var docs []internal.Document
	for _, d := range reply.Documents {
		doc := internal.Document{
			Content:   d.Content,
			Title:     d.Title,
			Author:    d.Author,
			Topic:     d.Topic,
			Watermark: d.Watermark,
		}
		docs = append(docs, doc)
	}
	return endpoint.GetResponse{Documents: docs, Err: reply.Err}, nil
}

func decodeGRPCAddResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.AddReply)
	return endpoint.AddResponse{TicketID: reply.TicketID, Err: reply.Err}, nil
}

func decodeGRPCUpdateResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.UpdateReply)
	return endpoint.UpdateResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func decodeGRPCRemoveResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.RemoveReply)
	return endpoint.RemoveResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func decodeGRPCServiceStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*db.ServiceStatusReply)
	return endpoint.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
