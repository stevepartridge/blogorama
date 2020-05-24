package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	blog "github.com/stevepartridge/blogorama"
	"github.com/stevepartridge/blogorama/pkg/log"
	pb "github.com/stevepartridge/blogorama/protos/go"
)

func (s *blogService) CreateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {
	log.Debug().Interface("user", req).Msg("CreateUser")

	if req == nil {
		return nil, ErrInvalidRequest
	}

	user := blog.UserFromProto(req)
	// user.CreatedByID = extractUserIDFromContext(ctx)

	err := blog.Store.CreateUser(&user)
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.UserResponse{
		User: user.ToProto(),
	}

	resp.User.ApiKey, err = blog.Store.GetAPIKeyForUserID(user.ID)
	if err != nil {
		return nil, handleError(err)
	}

	return resp, nil
}

func (s *blogService) UpdateUser(ctx context.Context, req *pb.User) (*pb.UserResponse, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	user := blog.UserFromProto(req)

	user.UpdatedByID = extractUserIDFromContext(ctx)

	err := blog.Store.UpdateUser(&user)
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.UserResponse{
		User: user.ToProto(),
	}

	return resp, nil
}

func (s *blogService) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	if req.User == nil {
		return nil, ErrInvalidRequest
	}

	user, err := blog.Store.GetUserByID(int(req.GetUser().Id))
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.UserResponse{
		User: user.ToProto(),
	}

	return resp, nil
}

func (s *blogService) GetUsers(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	list, info, err := blog.Store.GetUsers(int(req.Limit), int(req.Offset))
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.UserResponse{
		Users:  blog.UsersToProtos(list),
		Limit:  int32(info.Limit),
		Offset: int32(info.Offset),
		Total:  int32(info.Total),
	}

	return resp, nil
}

func (s *blogService) DeleteUser(ctx context.Context, req *pb.User) (*empty.Empty, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	err := blog.Store.DeleteUser(
		int(req.Id),
		extractUserIDFromContext(ctx),
	)

	if err != nil {
		return nil, handleError(err)
	}

	return &empty.Empty{}, nil
}
