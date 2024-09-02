// ApplicationServer\Comd_Open_Pos_Service\cmd\main.go
package main

import (
    "context"
    "log"
    "net"
    "os"
    "os/signal"
    "syscall"

    "github.com/SudarshanZone/MTF_OPEN_POS/config"
    "github.com/SudarshanZone/MTF_OPEN_POS/internal/repository"
    "github.com/SudarshanZone/MTF_OPEN_POS/internal/service"
    "google.golang.org/grpc"
    pb "github.com/SudarshanZone/MTF_OPEN_POS/generated"
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
    repo := repository.NewEquityRepository(db)

    // Initialize service
    srv := service.NewEquityService(repo)

    // Set up gRPC server
    lis, err := net.Listen("tcp", ":50054") 
    if err != nil {
        log.Fatalf("Failed to listen on port 50054: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterMtfPositionServiceServer(grpcServer,srv)

    // Handle graceful shutdown
    ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
    defer stop()

    go func() {
        <-ctx.Done()
        log.Println("Shutting down gRPC server...")
        grpcServer.GracefulStop()
    }()

    log.Println("Starting gRPC server on port 50054")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve gRPC server: %v", err)
    }
}
