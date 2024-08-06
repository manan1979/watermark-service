package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	pb "github.com/manan1979/watermark-service/api/pb/db"
	"github.com/manan1979/watermark-service/internal/database"
	dbsvc "github.com/manan1979/watermark-service/pkg/database"
	"gorm.io/gorm"

	"github.com/manan1979/watermark-service/pkg/database/endpoint"
	"github.com/manan1979/watermark-service/pkg/database/transport"
	"github.com/oklog/oklog/pkg/group"
	"google.golang.org/grpc"
)

const (
	defaultHTTPPort = "8081"
	defaultGRPCPort = "8082"
)

var (
	logger   log.Logger
	httpAddr = net.JoinHostPort("localhost", envString("HTTP_PORT", defaultHTTPPort))
	grpcAddr = net.JoinHostPort("localhost", envString("GRPC_PORT", defaultGRPCPort))
	db       *gorm.DB
)

func main() {
	var (
		service     = dbsvc.NewService()
		eps         = endpoint.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(eps)
		grpcServer  = transport.NewGRPCServer(eps)
	)

	var g group.Group
	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("tranasport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)

		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			httpListener.Close()
		})
	}

	{
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Log("transport", "gRPC", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "gRPC", "addr", grpcAddr)

			baseServer := grpc.NewServer(grpc.UnaryInterceptor(kitgrpc.Interceptor))
			pb.RegisterDatabaseServer(baseServer, grpcServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})

	}

	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("recieved signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}
	logger.Log("exit", g.Run())
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var err error
	db, err = database.Init(database.DefaultHost, database.DefaultPort, database.DefaultDBUser, database.DefaultDatabase, database.DefaultPassword)
	if err != nil {
		logger.Log("FATAL", fmt.Sprintf("failed to load db with error: %s", err.Error()))
		os.Exit(1)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Log("ERROR", fmt.Sprintf("failed to get sqlDB from gorm DB: %s", err.Error()))
		os.Exit(1)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			logger.Log("ERROR", fmt.Sprintf("failed to close the database connection: %s", err.Error()))
		}
	}()
}
