package trip_service_client

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
)

type tripServiceClient struct {
	client  pb.TripServiceClient
	url     string
	port    uint64
	version uint32
}

func NewTripServiceClient(url string, version uint32, port uint64) (TripServiceClient, error) {
	grpcClient, err := grpc.NewClient(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return &tripServiceClient{}, err
	}

	return &tripServiceClient{
		url:     url,
		version: version,
		port:    port,
		client:  pb.NewTripServiceClient(grpcClient),
	}, nil
}

func (us *tripServiceClient) SetVehicle(trip *pb.TripVehicle) (*pb.TripVehicleResponse, error) {
	response, err := us.client.SetVehicle(context.Background(), trip)

	if err != nil {
		return nil, err
	}

	if !response.Success {
		return nil, errors.New("error in set trip vehicle occurred")
	}

	return response, nil
}
