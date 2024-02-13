package repo

import (
	pb "USER_SERVICE/genproto/user-service"
)

// UserStorageI ...
type UserStorageI interface {
	Create(user *pb.User) (*pb.User, error)
	Update(user *pb.User) (*pb.User, error)
	Delete(request *pb.UserRequest) (*pb.User, error)
	Get(request *pb.UserRequest) (*pb.User, error)
	GetAll(request *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error)
}
