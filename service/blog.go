package service

import (
	"context"

	pb "github.com/stevepartridge/blogorama/protos/go"
)

type blogService struct{}

func (s *blogService) Info(c context.Context, req *pb.InfoRequest) (*pb.InfoResponse, error) {

	return &pb.InfoResponse{
		Name:    Name,
		BuiltAt: BuiltAt,
		Version: Version,
		Build:   Build,
		GitHash: GitHash,
	}, nil

}
