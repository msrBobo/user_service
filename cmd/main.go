package main

import (
	"USER_SERVICE/config"
	pb "USER_SERVICE/genproto/user-service"
	"USER_SERVICE/pkg/db"
	"USER_SERVICE/pkg/logger"
	"USER_SERVICE/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "USER_SERVICE")
	defer logger.Cleanup(log)

	log.Info("main: sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err, _ := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := service.NewUserService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
