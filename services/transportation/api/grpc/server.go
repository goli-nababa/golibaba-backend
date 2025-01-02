package grpc

import (
	"context"
	"transportation/app"
	"transportation/config"
	"transportation/internal/trip/domain"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
)

type tripServiceGRPCApi struct {
	pb.UnsafeTripServiceServer
	appContainer app.App
	config       config.Config
}

func NewTripServiceGRPCApi(appContainer app.App, cfg config.Config) pb.TripServiceServer {
	return &tripServiceGRPCApi{
		appContainer: appContainer,
		config:       cfg,
	}
}

func (t *tripServiceGRPCApi) SetVehicle(ctx context.Context, tripVehicle *pb.TripVehicle) (*pb.TripVehicleResponse, error) {
	s := t.appContainer.TripService(ctx)

	_, err := s.SetVehicle(ctx, domain.TripId(tripVehicle.TripId), domain.VehicleId(tripVehicle.VehicleId))
	if err != nil {
		return &pb.TripVehicleResponse{Success: false}, err
	}

	requests, err := s.GetVehicleRequests(ctx, domain.GetVehicleRequests{TripId: domain.TripId(tripVehicle.TripId)})

	if err != nil {
		return &pb.TripVehicleResponse{Success: false}, err
	}

	for _, v := range requests {
		s.DeleteVehicleRequest(ctx, v.ID)
	}

	return &pb.TripVehicleResponse{Success: true}, nil
}
