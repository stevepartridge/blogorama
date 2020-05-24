package service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"

	blog "github.com/stevepartridge/blogorama"
	pb "github.com/stevepartridge/blogorama/protos/go"
)

func (s *blogService) CreatePost(ctx context.Context, req *pb.Post) (*pb.PostResponse, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	post := blog.PostFromProto(req)
	post.CreatedByID = extractUserIDFromContext(ctx)

	err := blog.Store.CreatePost(&post)
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.PostResponse{
		Post: post.ToProto(),
	}

	return resp, nil
}

func (s *blogService) UpdatePost(ctx context.Context, req *pb.Post) (*pb.PostResponse, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	post := blog.PostFromProto(req)
	post.UpdatedByID = extractUserIDFromContext(ctx)

	err := blog.Store.UpdatePost(&post)
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.PostResponse{
		Post: post.ToProto(),
	}

	return resp, nil
}

func (s *blogService) GetPost(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	if req.Post == nil {
		return nil, ErrInvalidRequest
	}

	post, err := blog.Store.GetPostByID(int(req.GetPost().Id))
	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.PostResponse{
		Post: post.ToProto(),
	}

	return resp, nil
}

func (s *blogService) GetPosts(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	var (
		list []blog.Post
		info blog.ResultsInfo
		err  error
	)

	if req.UserId > 0 {
		list, info, err = blog.Store.GetPosts(int(req.UserId), int(req.Limit), int(req.Offset))
	} else {
		list, info, err = blog.Store.GetPosts(int(req.Limit), int(req.Offset))
	}

	if err != nil {
		return nil, handleError(err)
	}

	resp := &pb.PostResponse{
		Posts:  blog.PostsToProtos(list),
		Limit:  int32(info.Limit),
		Offset: int32(info.Offset),
		Total:  int32(info.Total),
	}

	return resp, nil
}

func (s *blogService) GetPostsByUser(ctx context.Context, req *pb.PostRequest) (*pb.PostResponse, error) {

	return s.GetPosts(ctx, req)
}

func (s *blogService) DeletePost(ctx context.Context, req *pb.Post) (*empty.Empty, error) {

	if req == nil {
		return nil, ErrInvalidRequest
	}

	return nil, nil
}
