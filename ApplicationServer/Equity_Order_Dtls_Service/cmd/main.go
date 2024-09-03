package main

import (
	"log"
	"net"

	"github.com/SudarshanZone/Equity_Ord_Dtls/config"
	pb "github.com/SudarshanZone/Equity_Ord_Dtls/generated"
	"github.com/SudarshanZone/Equity_Ord_Dtls/internal/repository"
	"github.com/SudarshanZone/Equity_Ord_Dtls/internal/service"
	"google.golang.org/grpc"
)

func main() {
	serviceName := "main"
	fileName := "config/EnvConfig.ini"

	cm := &config.ConfigManager{}

	// Load PostgreSQL configuration
	dbConfig, err := cm.LoadPostgreSQLConfig(serviceName, fileName)
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Get database connection
	connectionStatus := cm.GetDatabaseConnection(serviceName, *dbConfig)
	if connectionStatus != 0 {
		log.Fatalf("Failed to connect to the database")
	}

	db := cm.GetDB(serviceName)

	// Initialize the repository with the database connection
	repo := &repository.OrderDetailsRepository{Db: db}

	// Initialize the service with the repository
	orderDetailsService := service.NewOrderDetailsService(*repo)

	// Set up gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderDetailsService)

	// Listen on a TCP port
	listener, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("Failed to listen on port 50056: %v", err)
	}

	// Start the gRPC server
	log.Println("Starting gRPC server on port 50056...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
