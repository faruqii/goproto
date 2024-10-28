package services

import (
	"context"

	"github.com/faruqii/goproto/internal/domain/entities"
	"github.com/faruqii/goproto/internal/domain/repositories"
	products "github.com/faruqii/goproto/proto/products"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServiceServer struct {
	products.UnimplementedProductServiceServer
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) *ProductServiceServer {
	return &ProductServiceServer{repo: repo}
}

func (s *ProductServiceServer) CreateProduct(ctx context.Context, req *products.CreateProductRequest) (*products.CreateProductResponse, error) {
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

	response := &products.CreateProductResponse{
		Message: "Product created successfully",
		Result: &products.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		},
	}
	return response, nil

}

func (s *ProductServiceServer) SearchProducts(ctx context.Context, req *products.SearchProductsRequest) (*products.SearchProductsResponse, error) {
	product, err := s.repo.SearchProducts(req.Query)
	if err != nil {
		return nil, status.Error(codes.Internal, "Failed to search products")
	}

	response := &products.SearchProductsResponse{
		Products: make([]*products.Product, len(product)),
	}
	for i, product := range product {
		response.Products[i] = &products.Product{
			Id:          product.Id,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}
	}

	return response, nil
}
