package main

import (
	"fmt"
	"log"
	"net"

	"github.com/faruqii/goproto/internal/config/database"
	"github.com/faruqii/goproto/internal/domain/repositories"
	"github.com/faruqii/goproto/internal/services"
	"github.com/faruqii/goproto/proto/products"
	"github.com/faruqii/goproto/proto/users"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	// Listen on a TCP port
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server instance
	grpcServer := grpc.NewServer()

	// Repository
	productRepo := repositories.NewProductRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Register the ProductService with the server
	productService := services.NewProductService(productRepo)
	products.RegisterProductServiceServer(grpcServer, productService)

	// Register the UserService with the server
	userService := services.NewUserService(userRepo)
	users.RegisterUserServiceServer(grpcServer, userService)

	log.Println("gRPC server is running on port :50051")

	// Start serving requests
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
