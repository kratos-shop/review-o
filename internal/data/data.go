package data

import (
	"context"
	"review-o/internal/conf"

	v1 "review-o/api/review/v1"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"github.com/hashicorp/consul/api"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewRegistry, NewGrpcClient, NewData, NewOperationRepo)

// Data .
type Data struct {
	// grpc client
	grpcClient v1.ReviewClient
	logger     *log.Helper
}

// NewData .
func NewData(client v1.ReviewClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		grpcClient: client,
		logger:     log.NewHelper(logger),
	}, cleanup, nil
}

func NewGrpcClient(reg *consul.Registry, logger log.Logger) v1.ReviewClient {
	endpoint := "discovery:///review-service"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(reg))
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
		panic(err)
	}
	return v1.NewReviewClient(conn)
}

func NewRegistry(conf *conf.Registry) *consul.Registry {
	cfg := api.DefaultConfig()
	cfg.Address = conf.Consul.Address
	cfg.Scheme = conf.Consul.Scheme
	client, err := api.NewClient(cfg)
	if err != nil {
		log.Fatalf("new consul client failed: %v", err)
		panic(err)
	}
	consulClient := consul.New(client)
	log.Infof("new consul registry: %+v", consulClient)
	return consulClient
}
