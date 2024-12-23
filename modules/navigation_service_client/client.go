package navigation_service_client

import (
	"context"
	"fmt"
	pb "github.com/goli-nababa/golibaba-backend/modules/navigation_service_client/proto/gen/go/location/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	conn          *grpc.ClientConn
	location      pb.LocationServiceClient
	route         pb.RouteServiceClient
	defaultCtx    context.Context
	defaultCancel context.CancelFunc
}

type Config struct {
	Address string
	Timeout time.Duration
}

func NewClient(cfg Config) (*Client, error) {
	if cfg.Timeout == 0 {
		cfg.Timeout = 30 * time.Second
	}

	conn, err := grpc.NewClient(cfg.Address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)

	return &Client{
		conn:          conn,
		location:      pb.NewLocationServiceClient(conn),
		route:         pb.NewRouteServiceClient(conn),
		defaultCtx:    ctx,
		defaultCancel: cancel,
	}, nil
}

func (c *Client) Close() error {
	c.defaultCancel()
	return c.conn.Close()
}

func (c *Client) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.Location, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.location.CreateLocation(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create location: %w", err)
	}
	return resp.Location, nil
}

func (c *Client) GetLocation(ctx context.Context, id uint64) (*pb.Location, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.location.GetLocation(ctx, &pb.GetLocationRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("failed to get location: %w", err)
	}
	return resp.Location, nil
}

func (c *Client) UpdateLocation(ctx context.Context, req *pb.UpdateLocationRequest) (*pb.Location, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.location.UpdateLocation(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update location: %w", err)
	}
	return resp.Location, nil
}

func (c *Client) DeleteLocation(ctx context.Context, id uint64) error {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	_, err := c.location.DeleteLocation(ctx, &pb.DeleteLocationRequest{Id: id})
	if err != nil {
		return fmt.Errorf("failed to delete location: %w", err)
	}
	return nil
}

func (c *Client) ListLocations(ctx context.Context, activeOnly bool) ([]*pb.Location, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.location.ListLocations(ctx, &pb.ListLocationsRequest{
		ActiveOnly: activeOnly,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list locations: %w", err)
	}
	return resp.Locations, nil
}

func (c *Client) CreateRoute(ctx context.Context, req *pb.CreateRouteRequest) (*pb.Route, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.CreateRoute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create route: %w", err)
	}
	return resp.Route, nil
}

func (c *Client) GetRoute(ctx context.Context, id uint64) (*pb.Route, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.GetRoute(ctx, &pb.GetRouteRequest{Id: id})
	if err != nil {
		return nil, fmt.Errorf("failed to get route: %w", err)
	}
	return resp.Route, nil
}

func (c *Client) UpdateRoute(ctx context.Context, req *pb.UpdateRouteRequest) (*pb.Route, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.UpdateRoute(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to update route: %w", err)
	}
	return resp.Route, nil
}

func (c *Client) DeleteRoute(ctx context.Context, id uint64) error {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	_, err := c.route.DeleteRoute(ctx, &pb.DeleteRouteRequest{Id: id})
	if err != nil {
		return fmt.Errorf("failed to delete route: %w", err)
	}
	return nil
}

func (c *Client) SearchRoutes(ctx context.Context, req *pb.SearchRoutesRequest) ([]*pb.Route, int32, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.SearchRoutes(ctx, req)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to search routes: %w", err)
	}
	return resp.Routes, resp.TotalCount, nil
}

func (c *Client) CalculateRouteDistance(ctx context.Context, fromID, toID uint64, vehicleType pb.VehicleType) (*pb.CalculateRouteResponse, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.CalculateRouteDistance(ctx, &pb.CalculateRouteRequest{
		FromId:      fromID,
		ToId:        toID,
		VehicleType: vehicleType,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to calculate route distance: %w", err)
	}
	return resp, nil
}

func (c *Client) GetAvailableRoutes(ctx context.Context, req *pb.GetAvailableRoutesRequest) ([]*pb.Route, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.GetAvailableRoutes(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get available routes: %w", err)
	}
	return resp.Routes, nil
}

func (c *Client) ValidateRoute(ctx context.Context, routeID uint64, vehicleType pb.VehicleType) (bool, string, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.ValidateRoute(ctx, &pb.ValidateRouteRequest{
		RouteId:     routeID,
		VehicleType: vehicleType,
	})
	if err != nil {
		return false, "", fmt.Errorf("failed to validate route: %w", err)
	}
	return resp.IsValid, resp.ErrorMessage, nil
}

func (c *Client) ValidateRouteForTour(ctx context.Context, req *pb.ValidateRouteForTourRequest) (*pb.ValidateRouteForTourResponse, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.ValidateRouteForTour(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to validate route for tour: %w", err)
	}
	return resp, nil
}

func (c *Client) GetNearbyLocations(ctx context.Context, req *pb.GetNearbyLocationsRequest) ([]*pb.Location, map[uint64]float64, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.GetNearbyLocations(ctx, req)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get nearby locations: %w", err)
	}
	return resp.Locations, resp.Distances, nil
}

func (c *Client) GetPopularRoutes(ctx context.Context, req *pb.GetPopularRoutesRequest) ([]*pb.PopularRoute, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	resp, err := c.route.GetPopularRoutes(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get popular routes: %w", err)
	}
	return resp.Routes, nil
}

func (c *Client) GetOptimalRoutes(ctx context.Context, fromID uint64, toID uint64, vehicleTypes []pb.VehicleType, maxAlternatives int32) (pb.RouteService_GetOptimalRoutesClient, error) {
	if ctx == nil {
		ctx = c.defaultCtx
	}

	req := &pb.GetOptimalRoutesRequest{
		FromId:              fromID,
		ToId:                toID,
		AllowedVehicleTypes: vehicleTypes,
		MaxAlternatives:     maxAlternatives,
	}

	return c.route.GetOptimalRoutes(ctx, req)
}
