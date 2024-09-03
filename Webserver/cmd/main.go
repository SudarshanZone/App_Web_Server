package main

import (
    "fmt"
    "os"
    "time"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
    "github.com/SudarshanZone/Web_Server/internal/config"
    "github.com/SudarshanZone/Web_Server/internal/grpc"
    "github.com/SudarshanZone/Web_Server/internal/handlers"
    "github.com/SudarshanZone/Web_Server/internal/utils"
)

func init() {
    gin.SetMode(gin.ReleaseMode)
    
    utils.Logger.SetFormatter(&logrus.JSONFormatter{})
    utils.Logger.SetOutput(os.Stdout)
    utils.Logger.SetLevel(logrus.InfoLevel)
}

func main() {
    logger := utils.Logger.WithFields(logrus.Fields{
        "component": "webserver",
    })

    // Load config
    cfg, err := config.LoadConfig("internal/config/EnvConfig.ini")
    if err != nil {
        logger.Fatalf("Failed to load config: %v", err)
    }

    // Create gRPC client
    grpcClient ,err := grpc.NewClient(cfg.GRPCServerAddress1,cfg.GRPCServerAddress2,cfg.GRPCServerAddress3,cfg.GRPCServerAddress4,cfg.GRPCServerAddress5,cfg.GRPCServerAddress6)
    if err != nil {
        logger.Fatalf("Failed to create gRPC client: %v", err)
    }
    defer grpcClient.Close()

    // Initialize gRPC adapters for Clients
    grpcAdapter := grpc.NewGrpcAdapter(
        grpcClient.FnoPositionServiceClient,
        grpcClient.OrderDetailsServiceClient,
        grpcClient.CommodityPositionServiceClient,
        grpcClient.MtfPositionServiceClient,
        grpcClient.EquityMainServiceClient,
        grpcClient.EquityOrdServiceClient,
    )

    // Create facade
    facade := handlers.NewFacade(grpcAdapter)
    router := gin.New() 

    router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf(
            `{"time":"%s","method":"%s","path":"%s","status":%d,"latency":%s,"client_ip":"%s","BodySize":"%d","ClientIP":"%s"}` + "\n",
            param.TimeStamp.Format(time.RFC3339),
            param.Method,
            param.Path,
            param.StatusCode,
            param.Latency,
            param.ClientIP,
            param.BodySize,
            param.ClientIP,
        )
    }))
    router.Use(gin.Recovery())

    // CORS configuration
    router.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, 
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

    // Setup routes
    handlers.SetupRoutes(router, facade, logger)

    // Start the web server on port 8080
    logger.Info("Starting web server on ", cfg.ServerAddress)
    if err := router.Run(cfg.ServerAddress); err != nil {
        logger.Fatalf("Failed to start server: %v", err)
    }
}

