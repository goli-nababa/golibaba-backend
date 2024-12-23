package grpc_server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	di "navigation_service/app"
	"navigation_service/internal/common/types"
	"navigation_service/internal/location/domain"
	RoutingDomain "navigation_service/internal/routing/domain"
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
		ID:        domain.LocationID(uint(req.Id)),
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

func (s *routeServer) CreateRoute(ctx context.Context, req *pb.CreateRouteRequest) (*pb.CreateRouteResponse, error) {
	route := &RoutingDomain.Routing{
		Code:         req.Code,
		FromID:       uint(req.FromId),
		ToID:         uint(req.ToId),
		VehicleTypes: convertToVehicleTypes(req.VehicleTypes),
		Active:       true,
	}
	route, err := RoutingDomain.NewRoute(
		req.Code,
		uint(req.FromId),
		uint(req.ToId), 0,
		convertToVehicleTypes(req.VehicleTypes),
	)

	if err != nil {
		return nil, err
	}

	if err := s.app.RoutingService(ctx).CreateRouting(ctx, route); err != nil {
		return nil, err
	}

	return &pb.CreateRouteResponse{
		Route: convertToProtoRoute(route),
	}, nil
}

func (s *routeServer) GetRoute(ctx context.Context, req *pb.GetRouteRequest) (*pb.GetRouteResponse, error) {
	route, err := s.app.RoutingService(ctx).GetRouting(ctx, uint(req.Id))
	if err != nil {
		return nil, err
	}

	if route == nil {
		return &pb.GetRouteResponse{}, nil
	}

	return &pb.GetRouteResponse{
		Route: convertToProtoRoute(route),
	}, nil
}

func (s *routeServer) UpdateRoute(ctx context.Context, req *pb.UpdateRouteRequest) (*pb.UpdateRouteResponse, error) {
	route := &RoutingDomain.Routing{
		ID:           RoutingDomain.RoutingID(uint(req.Id)),
		Code:         req.Code,
		FromID:       uint(req.FromId),
		ToID:         uint(req.ToId),
		VehicleTypes: convertToVehicleTypes(req.VehicleTypes),
		Active:       req.Active,
	}

	if err := s.app.RoutingService(ctx).UpdateRouting(ctx, route); err != nil {
		return nil, err
	}

	return &pb.UpdateRouteResponse{
		Route: convertToProtoRoute(route),
	}, nil
}

func (s *routeServer) DeleteRoute(ctx context.Context, req *pb.DeleteRouteRequest) (*pb.DeleteRouteResponse, error) {
	if err := s.app.RoutingService(ctx).DeleteRouting(ctx, uint(req.Id)); err != nil {
		return nil, err
	}
	return &pb.DeleteRouteResponse{}, nil
}

func (s *routeServer) SearchRoutes(ctx context.Context, req *pb.SearchRoutesRequest) (*pb.SearchRoutesResponse, error) {
	filter := RoutingDomain.RouteFilter{
		FromID:      uint(req.FromId),
		ToID:        uint(req.ToId),
		VehicleType: types.VehicleType(req.VehicleType.String()),
		ActiveOnly:  req.ActiveOnly,
		PageSize:    int(req.PageSize),
		PageNumber:  int(req.PageNumber),
	}

	routes, err := s.app.RoutingService(ctx).FindRouting(ctx, filter)
	if err != nil {
		return nil, err
	}

	protoRoutes := make([]*pb.Route, len(routes))
	for i, route := range routes {
		protoRoutes[i] = convertToProtoRoute(&route)
	}

	return &pb.SearchRoutesResponse{
		Routes:     protoRoutes,
		TotalCount: int32(len(routes)),
	}, nil
}

func (s *routeServer) CalculateRouteDistance(ctx context.Context, req *pb.CalculateRouteRequest) (*pb.CalculateRouteResponse, error) {
	details, err := s.app.RoutingService(ctx).CalculateDistance(ctx, req.FromId, req.ToId, types.VehicleType(req.VehicleType.String()))
	if err != nil {
		return nil, err
	}

	return &pb.CalculateRouteResponse{
		Distance:      details.Distance,
		EstimatedTime: int32(details.EstimatedTime),
		EstimatedCost: details.EstimatedCost,
	}, nil
}

func (s *routeServer) ValidateRoute(ctx context.Context, req *pb.ValidateRouteRequest) (*pb.ValidateRouteResponse, error) {
	err := s.app.RoutingService(ctx).ValidateRoutingForVehicleType(ctx, uint(req.RouteId), types.VehicleType(req.VehicleType.String()))
	if err != nil {
		return &pb.ValidateRouteResponse{
			IsValid:      false,
			ErrorMessage: err.Error(),
		}, nil
	}

	return &pb.ValidateRouteResponse{
		IsValid: true,
	}, nil
}

func (s *routeServer) GetNearbyLocations(ctx context.Context, req *pb.GetNearbyLocationsRequest) (*pb.GetNearbyLocationsResponse, error) {
	locations, distances, err := s.app.RoutingService(ctx).FindNearbyLocations(ctx, req.LocationId, req.RadiusKm, types.LocationType(req.LocationType.String()))
	if err != nil {
		return nil, err
	}

	protoLocations := make([]*pb.Location, len(locations))
	for i, loc := range locations {
		protoLocations[i] = convertToProtoLocation(&loc)
	}

	return &pb.GetNearbyLocationsResponse{
		Locations: protoLocations,
		Distances: distances,
	}, nil
}

func (s *routeServer) GetAvailableRoutes(ctx context.Context, req *pb.GetAvailableRoutesRequest) (*pb.GetAvailableRoutesResponse, error) {
	filter := RoutingDomain.RouteFilter{
		FromID:      uint(req.FromId),
		ToID:        uint(req.ToId),
		VehicleType: types.VehicleType(req.VehicleType.String()),
		PageSize:    1000,
		PageNumber:  1,
	}

	routes, err := s.app.RoutingService(ctx).FindRouting(ctx, filter)
	if err != nil {
		return nil, err
	}
	protoRoutes := make([]*pb.Route, len(routes))
	for i, route := range routes {
		protoRoutes[i] = convertToProtoRoute(&route)
	}

	return &pb.GetAvailableRoutesResponse{
		Routes:     protoRoutes,
		TotalCount: int32(len(routes)),
	}, err
}
func (s *routeServer) GetOptimalRoutes(req *pb.GetOptimalRoutesRequest, stream pb.RouteService_GetOptimalRoutesServer) error {
	allowedVehicleTypes := convertToVehicleTypes(req.AllowedVehicleTypes)
	routes, err := s.app.RoutingService(stream.Context()).FindOptimalRoutes(
		stream.Context(),
		req.FromId,
		req.ToId,
		allowedVehicleTypes,
	)
	if err != nil {
		return err
	}

	for _, route := range routes {
		if int32(len(routes)) > req.MaxAlternatives {
			break
		}

		err := stream.Send(&pb.OptimalRouteResponse{
			Route:                convertToProtoRoute(route.Route),
			EfficiencyScore:      route.EfficiencyScore,
			OptimizationCriteria: route.Criteria,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *routeServer) GetRouteStatistics(ctx context.Context, req *pb.GetRouteStatisticsRequest) (*pb.GetRouteStatisticsResponse, error) {
	filter := RoutingDomain.StatisticsFilter{
		StartTime:   req.StartTime.AsTime(),
		EndTime:     req.EndTime.AsTime(),
		VehicleType: types.VehicleType(req.VehicleType.String()),
	}

	stats, err := s.app.RoutingService(ctx).GetRouteStatistics(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Convert map[string]int to map[string]int32
	routesByVehicleType := make(map[string]int32)
	for k, v := range stats.RoutesByVehicleType {
		routesByVehicleType[k] = int32(v)
	}

	return &pb.GetRouteStatisticsResponse{
		TotalRoutes:         int32(stats.TotalRoutes),
		TotalDistance:       stats.TotalDistance,
		RoutesByVehicleType: routesByVehicleType,
		MostPopularRoutes:   convertPopularRoutesToProto(stats.PopularRoutes),
	}, nil
}
func convertPopularRoutesToProto(routes []RoutingDomain.PopularRoute) []*pb.PopularRoute {
	protoRoutes := make([]*pb.PopularRoute, len(routes))
	for i, route := range routes {
		protoRoutes[i] = &pb.PopularRoute{
			Route:         convertToProtoRoute(route.Route),
			UsageCount:    int32(route.UsageCount),
			AverageRating: route.AverageRating,
		}
	}
	return protoRoutes
}
func convertToProtoRoute(route *RoutingDomain.Routing) *pb.Route {
	return &pb.Route{
		Id:           uint64(route.ID),
		Uuid:         route.UUID,
		Code:         route.Code,
		FromId:       uint64(route.FromID),
		ToId:         uint64(route.ToID),
		Distance:     route.Distance,
		VehicleTypes: convertToProtoVehicleTypes(route.VehicleTypes),
		Active:       route.Active,
		CreatedAt:    route.CreatedAt.Unix(),
		UpdatedAt:    route.UpdatedAt.Unix(),
	}
}

func convertToVehicleTypes(pbTypes []pb.VehicleType) []types.VehicleType {
	result := make([]types.VehicleType, len(pbTypes))
	for i, t := range pbTypes {
		result[i] = types.VehicleType(t.String())
	}
	return result
}

func convertToProtoVehicleTypes(types []types.VehicleType) []pb.VehicleType {
	result := make([]pb.VehicleType, len(types))
	for i, t := range types {
		result[i] = pb.VehicleType(pb.VehicleType_value[string(t)])
	}
	return result
}
