package grpc

import (
	"context"
	"fmt"
	"github.com/sergeygardner/meal-planner-api/application/handler"
	protoBuf "github.com/sergeygardner/meal-planner-api/ui/grpc/model"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	protoBuf.UnimplementedAuthServer
}

func (s *server) Credentials(_ context.Context, authCredentialsDTO *protoBuf.AuthCredentialsDTO) (*protoBuf.AuthConfirmation, error) {
	authConfirmation, _, errorAuthCredentials := handler.AuthCredentials(authCredentialsDTO.UserCredentialsDTO)

	if errorAuthCredentials != nil {
		return nil, errorAuthCredentials
	}

	return &protoBuf.AuthConfirmation{AuthConfirmation: *authConfirmation}, nil
}

func GetServer(port int) (*grpc.Server, net.Listener) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	protoBuf.RegisterAuthServer(grpcServer, &server{})

	return grpcServer, listener
}
