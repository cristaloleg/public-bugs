package main

import (
	"context"
	"log"
	"net"

	"github.com/cristaloleg/public-bugs/protobuf1135/model"
	"github.com/cristaloleg/public-bugs/protobuf1135/svc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func main() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()

	svc.RegisterStatusServiceServer(s, &Server{})
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := svc.NewStatusServiceClient(conn)

	req := &svc.Request{
		Statuses: []model.Status{},
	}

	_, err = client.GetStatues(ctx, req)
	if err != nil {
		log.Fatalf("GetLeads failed: %v", err)
	}
}

var _ svc.StatusServiceServer = (*Server)(nil)

type Server struct {
}

func (*Server) GetStatues(context.Context, *svc.Request) (*svc.Response, error) {
	return nil, nil
}
