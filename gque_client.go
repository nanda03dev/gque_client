package gqueclient

import (
	"log"

	pb "github.com/nanda03dev/gque_client/proto"
	"google.golang.org/grpc"
)

type Client struct {
	URI        string // Ex: http://localhost:5456
	AppName    string
	GqueClient pb.GqueServiceClient
}

// Create new Gque client,
// URI string  Ex: http://localhost:5454
func Connect(URI string, appName string) *Client {
	var client = &Client{
		URI: URI,
	}

	conn, err := grpc.Dial(URI, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	} else {
		log.Println("conected to Gque - gRPC Server")
	}

	client.GqueClient = pb.NewGqueServiceClient(conn)
	client.AppName = appName

	return client
}
