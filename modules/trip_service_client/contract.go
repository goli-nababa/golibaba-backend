package trip_service_client

import (
	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
)

type TripServiceClient interface {
	SetVehicle(tripVehicle *pb.TripVehicle) (*pb.TripVehicleResponse, error)
}
