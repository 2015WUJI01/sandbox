package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"main/ep2.server_and_client/pb"
	"time"
)

func main() {

	conn, err := grpc.Dial("localhost:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewTrackingServerClient(conn)
	SendiRecord(client, &pb.IncomingRecord{
		Id:              123,
		VisitTime:       timestamppb.New(time.Now()),
		Ip:              "127.0.0.1",
		HttpReferer:     "",
		HttpUserAgent:   "",
		RequestUri:      "",
		TrafficType:     0,
		RetentionUserId: "",
		Source:          "",
		SourceGroup:     "",
		Country:         "",
		RefKeyword:      "",
		ServerId:        "",
		CurrServerId:    "",
	})
}

func SendiRecord(client pb.TrackingServerClient, iRecord *pb.IncomingRecord) {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	_, err := client.AddIncomingRecord(context.Background(), iRecord)
	if err != nil {
		log.Fatalf("failed: %v", err)
	}
	// log.Println(resp)
}
