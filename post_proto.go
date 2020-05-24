package blog

import (
	"time"

	pb "github.com/stevepartridge/blogorama/protos/go"
)

// ToProto converts Post to the protobuf pb.Post
func (p *Post) ToProto() *pb.Post {
	proto := &pb.Post{
		Id:          int32(p.ID),
		Title:       p.Title,
		Content:     p.Content,
		Private:     p.Private,
		UpdatedById: int32(p.UpdatedByID),
		CreatedById: int32(p.CreatedByID),
	}

	if !p.UpdatedAt.IsZero() {
		proto.UpdatedAt = p.UpdatedAt.Format(time.RFC3339Nano)
	}

	if !p.CreatedAt.IsZero() {
		proto.CreatedAt = p.CreatedAt.Format(time.RFC3339Nano)
	}

	return proto
}

// PostFromProto converts a protobuf pb.Post to Post
func PostFromProto(proto *pb.Post) Post {
	if proto == nil {
		return Post{}
	}

	post := Post{
		ID:      int(proto.GetId()),
		Title:   proto.GetTitle(),
		Content: proto.GetContent(),
		Private: proto.GetPrivate(),
	}

	return post
}

// PostsToProtos converts a list of Post/s to a list of protobuf pb.Post/s
func PostsToProtos(list []Post) []*pb.Post {
	result := make([]*pb.Post, len(list))
	for i := range list {
		result[i] = list[i].ToProto()
	}
	return result
}
