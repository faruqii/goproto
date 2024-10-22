package services

import (
	"context"

	"github.com/faruqii/goproto/internal/domain/entities"
	"github.com/faruqii/goproto/internal/domain/repositories"
	"github.com/faruqii/goproto/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServiceServer struct {
	proto.UnimplementedProductServiceServer
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductServiceServer {
	return &ProductServiceServer{repo: repo}
}

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *proto.CreateProductRequest) (*proto.CreateProductResponse, error) {
	if req.Name == "" || req.Price <= 0 {
		return nil, status.Error(codes.InvalidArgument, "Invalid product details")
	}
	product := &entities.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}

	if err := s.repo.CreateProduct(product); err != nil {
		return nil, status.Error(codes.InvalidArgument, "Failed to create product")
	}

	response := &proto.CreateProductResponse{
		Message: "Product created successfully",
		Result: &proto.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}

	return response, nil

}
