package services

import (
	"context"

	"github.com/faruqii/goproto/internal/domain/entities"
	"github.com/faruqii/goproto/internal/domain/repositories"
	"github.com/faruqii/goproto/proto/users"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceServer struct {
	users.UnimplementedUserServiceServer
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserServiceServer {
	return &UserServiceServer{repo: repo}
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *users.CreateUserRequest) (*users.CreateUserResponse, error) {
	if req.Name == "" || req.Email == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid user details")
	}

	// hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to hash password")
	}

	user := &entities.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPwd),
	}

	if err := s.repo.CreateUser(user); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to create user")
	}

	response := &users.CreateUserResponse{
		Message: "User created successfully",
		Result: &users.UserResponse{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return response, nil
}
