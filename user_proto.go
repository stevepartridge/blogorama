package blog

import (
	"time"

	pb "github.com/stevepartridge/blogorama/protos/go"
)

// ToProto converts User to the protobuf pb.User
func (u *User) ToProto() *pb.User {
	proto := &pb.User{
		Id:          int32(u.ID),
		Name:        u.Name,
		Active:      u.Active,
		UpdatedById: int32(u.UpdatedByID),
		CreatedById: int32(u.CreatedByID),
	}

	if !u.UpdatedAt.IsZero() {
		proto.UpdatedAt = u.UpdatedAt.Format(time.RFC3339Nano)
	}

	if !u.CreatedAt.IsZero() {
		proto.CreatedAt = u.CreatedAt.Format(time.RFC3339Nano)
	}

	return proto
}

// UserFromProto converts a protobuf pb.User to User
func UserFromProto(proto *pb.User) User {
	if proto == nil {
		return User{}
	}

	user := User{
		ID:     int(proto.GetId()),
		Name:   proto.GetName(),
		Active: proto.GetActive(),
	}

	return user
}

// UsersToProtos converts a list of User/s to a list of protobuf pb.User/s
func UsersToProtos(list []User) []*pb.User {
	result := make([]*pb.User, len(list))
	for i := range list {
		result[i] = list[i].ToProto()
	}
	return result
}
