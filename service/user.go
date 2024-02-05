package service

import (
	pb "USER_SERVICE/genproto/user-service"
	l "USER_SERVICE/pkg/logger"
	"USER_SERVICE/storage"
	"context"

	"github.com/jmoiron/sqlx"
)

// UserService ...
type UserService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedUserServiceServer
}

// NewUserService ...
func NewUserService(db *sqlx.DB, log l.Logger) *UserService {
	return &UserService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *UserService) Create(ctx context.Context, req *pb.User) (*pb.User, error) {
	user, err := s.storage.User().Create(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Update(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, err := s.storage.User().Update(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Delete(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, err := s.storage.User().Delete(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Get(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	user, err := s.storage.User().Get(req)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAll(ctx context.Context, req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := s.storage.User().GetAll(req)
	if err != nil {
		return nil, err
	}
	return users, nil
}
