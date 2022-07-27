package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"main/ep2.server_and_client/pb"
	"net"
)

type TrackingServer struct {
	pb.TrackingServerServer
}

func (s *TrackingServer) AddIncomingRecord(ctx context.Context, iRecord *pb.IncomingRecord) (*pb.Response, error) {
	fmt.Printf("ID:%d Time:%s", iRecord.Id,
		iRecord.VisitTime.AsTime().Format("2006-01-02 15:04:05"))
	return &pb.Response{}, nil
}

func (s *TrackingServer) AddOutgoingRecord(ctx context.Context, oRecord *pb.OutgoingRecord) (*pb.Response, error) {
	return nil, nil
}

func (s *TrackingServer) AddPagevisitRecord(ctx context.Context, pRecord *pb.PagevisitRecord) (*pb.Response, error) {
	return nil, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterTrackingServerServer(grpcServer, new(TrackingServer))
	lis, _ := net.Listen("tcp", "localhost:8090")
	_ = grpcServer.Serve(lis)
}
