package grpc

import (
	build "google.golang.org/genproto/googleapis/devtools/build/v1"
	"google.golang.org/grpc"

	"github.com/buildbarn/bb-portal/ent/gen/ent"
	"github.com/buildbarn/bb-portal/internal/api/grpc/bes"
	"github.com/buildbarn/bb-portal/pkg/processing"
)

type Server = grpc.Server

func NewServer(db *ent.Client, blobArchiver processing.BlobMultiArchiver, opts ...grpc.ServerOption) *grpc.Server {
	grpcServer := grpc.NewServer(opts...)

	build.RegisterPublishBuildEventServer(grpcServer, bes.New(db, blobArchiver))
	return grpcServer
}
