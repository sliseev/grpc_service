package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sliseev/grpc_service/api"
	"google.golang.org/grpc"
)

type Feature struct {
	name      string
	latitude  int
	longitude int
}

type grpcServer struct {
	pb.UnimplementedGrpcServiceServer
	features []Feature
}

func (s *grpcServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	for _, feature := range s.features {
		if feature.latitude == int(point.Latitude) && feature.longitude == int(point.Longitude) {
			return &pb.Feature{Location: point, Name: feature.name}, nil
		}
	}
	return &pb.Feature{Location: point}, nil
}

func (s *grpcServer) ListFeatures(rect *pb.Rectangle, stream pb.GrpcService_ListFeaturesServer) error {
	for _, feature := range s.features {
		f := &pb.Feature{
			Location: &pb.Point{
				Longitude: int32(feature.longitude),
				Latitude:  int32(feature.latitude),
			},
			Name: feature.name,
		}
		if err := stream.Send(f); err != nil {
			return err
		}
	}
	return nil
}

func (s *grpcServer) RecordRoute(stream pb.GrpcService_RecordRouteServer) error {
	return nil
}

func (s *grpcServer) RouteChat(stream pb.GrpcService_RouteChatServer) error {
	return nil
}

func newServer() *grpcServer {
	s := &grpcServer{features: []Feature{
		{
			name:      "Taxi1",
			latitude:  1,
			longitude: 1,
		},
		{
			name:      "Taxi2",
			latitude:  2,
			longitude: 2,
		},
		{
			name:      "Taxi3",
			latitude:  3,
			longitude: 3,
		},
	}}
	return s
}

func main() {
	lis, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterGrpcServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
