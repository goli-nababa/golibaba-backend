package trip_service_client

import (
	"fmt"
	"testing"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
)

func TestSetVehicle(t *testing.T) {
	client, err := NewTripServiceClient("localhost", 1, 8082)
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	response, err := client.SetVehicle(&pb.TripVehicle{TripId: 2, VehicleId: 3, VehicleSpeed: 50})

	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	fmt.Println(response)
}
