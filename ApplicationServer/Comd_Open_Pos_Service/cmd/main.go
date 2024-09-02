// ApplicationServer\Comd_Open_Pos_Service\cmd\main.go
package main

import (
    "context"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "github.com/SudarshanZone/Commo_Open_Pos/config"
    "github.com/SudarshanZone/Commo_Open_Pos/internal/repository"
    "github.com/SudarshanZone/Commo_Open_Pos/internal/service"
    "google.golang.org/grpc"
    pb "github.com/SudarshanZone/Commo_Open_Pos/generated"
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
    repo := repository.NewCCPRepository(db)

    // Initialize service
    srv := service.NewCCPService(repo)

    // Set up gRPC server
    lis, err := net.Listen("tcp", ":50053") 
    if err != nil {
        log.Fatalf("Failed to listen on port 50053: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterCCPServiceServer(grpcServer, srv)

    // Handle graceful shutdown
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    go func() {
        <-ctx.Done()
        log.Println("Shutting down gRPC server...")
        grpcServer.GracefulStop()
    }()

    log.Println("Starting gRPC server on port 50053")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}
