package main

import (
	"context"
	"flag"
	"github.com/sergeygardner/meal-planner-api/domain/dto"
	log "github.com/sirupsen/logrus"
	"time"

	protoBuf "github.com/sergeygardner/meal-planner-api/ui/grpc/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Could not close the connection: %s", err.Error())
		}
	}(conn)
	c := protoBuf.NewAuthClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Credentials(ctx, &protoBuf.AuthCredentialsDTO{UserCredentialsDTO: dto.UserCredentialsDTO{Username: "username", Password: "password"}})

	if err != nil {
		log.Fatalf("could not send the authentication request: %v", err)
	}

	log.Printf("It has sent the authentication request: %s", r.GetMessage())
}
