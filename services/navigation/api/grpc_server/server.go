package grpc_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	di "navigation_service/app"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	pb "navigation_service/proto/gen/go/location/v1"
	"net"
)

type locationServer struct {
	pb.UnimplementedLocationServiceServer
	app di.App
}

type routeServer struct {
	pb.UnimplementedRouteServiceServer
	app di.App
}

func newServer(app di.App) (*locationServer, *routeServer) {
	return &locationServer{app: app}, &routeServer{app: app}
}

func Bootstrap(app di.App, port uint) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	locationSrv, routeSrv := newServer(app)

	pb.RegisterLocationServiceServer(s, locationSrv)
	pb.RegisterRouteServiceServer(s, routeSrv)

	reflection.Register(s)

	fmt.Printf("gRPC server listening on port %d\n", port)
	if err := s.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}

	return nil
}

func (s *locationServer) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error) {
	location := &domain.Location{
		Name:      req.Name,
		Type:      types.LocationType(req.Type.String()),
		Address:   req.Address,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Active:    true,
	}

	if err := s.app.LocationService(ctx).CreateLocation(ctx, location); err != nil {
		return nil, err
	}

	return &pb.CreateLocationResponse{
		Location: convertToProtoLocation(location),
	}, nil
}

func (s *locationServer) GetLocation(ctx context.Context, req *pb.GetLocationRequest) (*pb.GetLocationResponse, error) {
	location, err := s.app.LocationService(ctx).GetLocation(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	if location == nil {
		return &pb.GetLocationResponse{}, nil
	}

	return &pb.GetLocationResponse{
		Location: convertToProtoLocation(location),
	}, nil
}

func (s *locationServer) UpdateLocation(ctx context.Context, req *pb.UpdateLocationRequest) (*pb.UpdateLocationResponse, error) {
	location := &domain.Location{
		ID:        uint(req.Id),
		Name:      req.Name,
		Type:      types.LocationType(req.Type.String()),
		Address:   req.Address,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Active:    req.Active,
	}

	if err := s.app.LocationService(ctx).UpdateLocation(ctx, location); err != nil {
		return nil, err
	}

	return &pb.UpdateLocationResponse{
		Location: convertToProtoLocation(location),
	}, nil
}

func (s *locationServer) DeleteLocation(ctx context.Context, req *pb.DeleteLocationRequest) (*pb.DeleteLocationResponse, error) {
	if err := s.app.LocationService(ctx).DeleteLocation(ctx, uint(req.Id)); err != nil {
		return nil, err
	}

	return &pb.DeleteLocationResponse{}, nil
}

func (s *locationServer) ListLocations(ctx context.Context, req *pb.ListLocationsRequest) (*pb.ListLocationsResponse, error) {
	locations, err := s.app.LocationService(ctx).ListLocations(ctx, req.ActiveOnly)
	if err != nil {
		return nil, err
	}

	protoLocations := make([]*pb.Location, len(locations))
	for i, loc := range locations {
		protoLocations[i] = convertToProtoLocation(&loc)
	}

	return &pb.ListLocationsResponse{
		Locations:  protoLocations,
		TotalCount: int32(len(locations)),
	}, nil
}

func convertToProtoLocation(location *domain.Location) *pb.Location {
	return &pb.Location{
		Id:        uint64(location.ID),
		Name:      location.Name,
		Type:      pb.LocationType(pb.LocationType_value[string(location.Type)]),
		Address:   location.Address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Active:    location.Active,
		CreatedAt: location.CreatedAt.Unix(),
		UpdatedAt: location.UpdatedAt.Unix(),
	}
}
