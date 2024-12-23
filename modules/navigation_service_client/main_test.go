package navigation_service_client

import (
	"context"
	"fmt"
	pb "github.com/goli-nababa/golibaba-backend/modules/navigation_service_client/proto/gen/go/location/v1"
	"github.com/stretchr/testify/require"
	"log"
	"strconv"
	"testing"
	"time"
)

const (
	serviceAddress = "localhost:8083"
	timeout        = 30 * time.Second
)

func setupIntegrationTest(t *testing.T) *Client {
	client, err := NewClient(Config{
		Address: serviceAddress,
		Timeout: timeout,
	})
	require.NoError(t, err)
	require.NotNil(t, client)
	return client
}

func TestIntegration_LocationLifecycle(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	defer func(client *Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	ctx := context.Background()

	loc, err := client.CreateLocation(ctx, &pb.CreateLocationRequest{
		Name:      "Tehran International Airport",
		Type:      pb.LocationType_AIRPORT,
		Address:   "Tehran Province, Tehran, Terminal 1",
		Latitude:  35.416111,
		Longitude: 51.152222,
	})
	require.NoError(t, err)
	require.NotNil(t, loc)
	require.NotZero(t, loc.Id)
	require.Equal(t, "Tehran International Airport", loc.Name)

	fetchedLoc, err := client.GetLocation(ctx, loc.Id)
	require.NoError(t, err)
	require.NotNil(t, fetchedLoc)
	require.Equal(t, loc.Id, fetchedLoc.Id)
	require.Equal(t, loc.Name, fetchedLoc.Name)

	updatedLoc, err := client.UpdateLocation(ctx, &pb.UpdateLocationRequest{
		Id:        loc.Id,
		Name:      "IKA International Airport",
		Type:      pb.LocationType_AIRPORT,
		Address:   loc.Address,
		Latitude:  loc.Latitude,
		Longitude: loc.Longitude,
		Active:    true,
	})
	require.NoError(t, err)
	require.NotNil(t, updatedLoc)
	require.Equal(t, "IKA International Airport", updatedLoc.Name)

	locations, err := client.ListLocations(ctx, true)
	require.NoError(t, err)
	require.NotEmpty(t, locations)
	found := false
	for _, l := range locations {
		if l.Id == loc.Id {
			found = true
			break
		}
	}
	require.True(t, found)

	err = client.DeleteLocation(ctx, loc.Id)
	require.NoError(t, err)

	resp, err := client.GetLocation(ctx, loc.Id)

	require.Nil(t, resp)
}

func TestIntegration_RouteLifecycle(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	defer func(client *Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	ctx := context.Background()

	sourceLoc, err := client.CreateLocation(ctx, &pb.CreateLocationRequest{
		Name:      "Tehran Airport",
		Type:      pb.LocationType_AIRPORT,
		Address:   "Tehran",
		Latitude:  35.416111,
		Longitude: 51.152222,
	})
	require.NoError(t, err)

	destLoc, err := client.CreateLocation(ctx, &pb.CreateLocationRequest{
		Name:      "Mashhad Airport",
		Type:      pb.LocationType_AIRPORT,
		Address:   "Mashhad",
		Latitude:  36.234,
		Longitude: 59.643,
	})
	require.NoError(t, err)

	code := strconv.Itoa(int(time.Now().Unix()))

	route, err := client.CreateRoute(ctx, &pb.CreateRouteRequest{
		Code:   code,
		FromId: sourceLoc.Id,
		ToId:   destLoc.Id,
		VehicleTypes: []pb.VehicleType{
			pb.VehicleType_AIRPLANE,
		},
	})
	require.NoError(t, err)
	require.NotNil(t, route)
	require.NotZero(t, route.Id)

	distance, err := client.CalculateRouteDistance(ctx, sourceLoc.Id, destLoc.Id, pb.VehicleType_AIRPLANE)
	require.NoError(t, err)
	require.NotNil(t, distance)
	require.True(t, distance.Distance > 0)
	require.True(t, distance.EstimatedTime > 0)

	routes, count, err := client.SearchRoutes(ctx, &pb.SearchRoutesRequest{
		FromId:      sourceLoc.Id,
		ToId:        destLoc.Id,
		VehicleType: pb.VehicleType_AIRPLANE,
		ActiveOnly:  true,
		PageSize:    10,
		PageNumber:  1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, routes)
	require.True(t, count > 0)

	stream, err := client.GetOptimalRoutes(ctx, sourceLoc.Id, destLoc.Id,
		[]pb.VehicleType{pb.VehicleType_AIRPLANE}, 5)
	require.NoError(t, err)

	optimalRouteCount := 0
	for {
		route, err := stream.Recv()
		if err != nil {
			break
		}
		require.NotNil(t, route)
		require.NotNil(t, route.Route)
		require.True(t, route.EfficiencyScore > 0)
		optimalRouteCount++
	}
	require.True(t, optimalRouteCount > 0)

	nearby, distances, err := client.GetNearbyLocations(ctx, &pb.GetNearbyLocationsRequest{
		LocationId:   sourceLoc.Id,
		RadiusKm:     1000,
		LocationType: pb.LocationType_AIRPORT,
	})
	require.NoError(t, err)
	require.NotEmpty(t, nearby)
	require.NotEmpty(t, distances)

	err = client.DeleteRoute(ctx, route.Id)
	require.NoError(t, err)

	err = client.DeleteLocation(ctx, sourceLoc.Id)
	require.NoError(t, err)

	err = client.DeleteLocation(ctx, destLoc.Id)
	require.NoError(t, err)
}

func TestIntegration_ConcurrentOperations(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	defer func(client *Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	ctx := context.Background()

	numLocations := 5
	errChan := make(chan error, numLocations)
	locChan := make(chan *pb.Location, numLocations)

	for i := 0; i < numLocations; i++ {
		go func(index int) {
			loc, err := client.CreateLocation(ctx, &pb.CreateLocationRequest{
				Name:      fmt.Sprintf("Test Location %d", index),
				Type:      pb.LocationType_AIRPORT,
				Address:   fmt.Sprintf("Address %d", index),
				Latitude:  35.0 + float64(index),
				Longitude: 51.0 + float64(index),
			})
			if err != nil {
				errChan <- err
				return
			}
			locChan <- loc
		}(i)
	}

	locations := make([]*pb.Location, 0, numLocations)
	for i := 0; i < numLocations; i++ {
		select {
		case err := <-errChan:
			require.NoError(t, err)
		case loc := <-locChan:
			locations = append(locations, loc)
		case <-time.After(30 * time.Second):
			t.Fatal("Timeout waiting for concurrent operations")
		}
	}

	require.Len(t, locations, numLocations)

	for _, loc := range locations {
		err := client.DeleteLocation(ctx, loc.Id)
		require.NoError(t, err)
	}
}

func TestIntegration_ErrorCases(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	client := setupIntegrationTest(t)
	defer func(client *Client) {
		err := client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(client)

	ctx := context.Background()

	// Test invalid location ID
	data, err := client.GetLocation(ctx, 99999999)
	require.Nil(t, data)

	_, err = client.CreateLocation(ctx, &pb.CreateLocationRequest{
		Name:      "Invalid Location",
		Type:      pb.LocationType_AIRPORT,
		Address:   "Test",
		Latitude:  200, // Invalid latitude
		Longitude: 51.0,
	})
	require.Error(t, err)

	// Test timeout
	shortCtx, cancel := context.WithTimeout(ctx, 1*time.Microsecond)
	defer cancel()
	time.Sleep(time.Millisecond)
	_, err = client.ListLocations(shortCtx, true)
	require.Error(t, err)
}
