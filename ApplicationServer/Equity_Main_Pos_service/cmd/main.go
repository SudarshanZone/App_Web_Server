package main

import (
    "context"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "github.com/SudarshanZone/Equity_Main_Open_Pos/config"
    "github.com/SudarshanZone/Equity_Main_Open_Pos/internal/repository"
    "github.com/SudarshanZone/Equity_Main_Open_Pos/internal/service"
    "google.golang.org/grpc"
    pb "github.com/SudarshanZone/Equity_Main_Open_Pos/generated"
)

func main() {
    serviceName := "main"
    fileName := "config/EnvConfig.ini"

    // Initialize ConfigManager
    cm := &config.ConfigManager{}

    // Load PostgreSQL config
    dbConfig, err := cm.LoadPostgreSQLConfig(serviceName, fileName)
    if err != nil {
        log.Fatalf("Failed to load database config: %v", err)
    }

    // Get database connection
    connectionStatus := cm.GetDatabaseConnection(serviceName, *dbConfig)
    if connectionStatus != 0 {
        log.Fatalf("Failed to connect to the database with status code %d", connectionStatus)
    }

    // Retrieve the database instance
    db := cm.GetDB(serviceName)

    // Initialize repository with concrete implementation
    repo := repository.NewPositionRepository(db)

    // Initialize service
    srv := service.NewPositionService(repo)

    // Set up gRPC server
    lis, err := net.Listen("tcp", ":50055")
    if err != nil {
        log.Fatalf("Failed to listen on port 50055: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterPositionServiceServer(grpcServer, srv)

    // Handle graceful shutdown
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    go func() {
        <-ctx.Done()
        log.Println("Shutting down gRPC server...")
        grpcServer.GracefulStop()
    }()

    log.Println("Starting gRPC server on port 50055")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}
