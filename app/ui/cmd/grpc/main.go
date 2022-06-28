package main

import (
	"flag"
	InfrastructureService "github.com/sergeygardner/meal-planner-api/infrastructure/service"
	"github.com/sergeygardner/meal-planner-api/ui/grpc"
	log "github.com/sirupsen/logrus"
)

var (
	flagDev      bool
	flagGRPCPort int
)

func init() {
	flag.BoolVar(&flagDev, "dev", false, "To use Dev")
	flag.IntVar(&flagGRPCPort, "gRPCPort", 50051, "To use a port in gRPC")
}

func main() {
	flag.Parse()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetReportCaller(true)

	if flagDev {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}

	InfrastructureService.PrepareCache()
	InfrastructureService.PreparePersistence()

	server, listener := grpc.GetServer(flagGRPCPort)
	log.Printf("server listening at %v", listener.Addr())
	log.Fatalf("failed to serve: %v", server.Serve(listener))
}
