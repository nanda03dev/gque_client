package gqueclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/nanda03dev/gque_client/proto"
)

func (gqueClient *Client) CreateQueue(queue Queue) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = gqueClient.GqueClient

	requestBody := &pb.QueueCreateRequest{
		QueueName: queue.Name,
		Time:      queue.Time,
	}

	return gRPC.CreateQueue(ctx, requestBody)
}

func (gqueClient *Client) CreateBroadcast(broadcast Broadcast) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = gqueClient.GqueClient

	requestBody := &pb.BroadcastCreateRequest{
		BroadcastName: broadcast.Name,
		QueueNames:    broadcast.QueueNames,
	}

	return gRPC.CreateBroadcast(ctx, requestBody)
}

func (gqueClient *Client) PushMessage(queueMessage QueueMessageType) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = gqueClient.GqueClient

	messageString, MarshallErr := json.Marshal(queueMessage.Data)

	if MarshallErr != nil {
		return nil, errors.New(ERROR_WHILE_MARSHAL_JSON)
	}

	requestBody := &pb.PushMessageRequest{
		QueueName: queueMessage.Name,
		Message:   string(messageString),
	}

	return gRPC.PushMessage(ctx, requestBody)
}

func (gqueClient *Client) BroadcastMessage(broadcastMessage BroadcastMessageType) (*pb.SuccessResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var gRPC = gqueClient.GqueClient

	messageString, MarshallErr := json.Marshal(broadcastMessage.Data)

	if MarshallErr != nil {
		return nil, errors.New(ERROR_WHILE_MARSHAL_JSON)
	}

	requestBody := &pb.BroadcastMessageRequest{
		BroadcastName: broadcastMessage.Name,
		Message:       string(messageString),
	}

	return gRPC.BroadcastMessage(ctx, requestBody)
}

func (gqueClient *Client) Consume(ConsumerRequest ConsumerRequestType, receiveChan chan MessageType) error {

	ctx := context.WithoutCancel(context.Background())

	var gRPC = gqueClient.GqueClient

	requestBody := &pb.ConsumerRequest{
		QueueName: ConsumerRequest.QueueName,
	}

	stream, err := gRPC.ConsumeQueueMessages(ctx, requestBody)
	if err != nil {
		return err
	}

	go func() {
		defer close(receiveChan)
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				// The server has finished sending messages
				fmt.Println(" \n The server has finished sending messages:")

				break
			}
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				return
			}

			var newMessage MessageType

			var UnMarshallErr = json.Unmarshal([]byte(msg.Message), &newMessage)

			if UnMarshallErr != nil {
				log.Printf(" Error while UnMarshall message: %v", UnMarshallErr)
				log.Printf(" skipping this message : %v", msg.Message)
			}
			// Send the received message to the channel
			receiveChan <- newMessage
		}
	}()

	return nil
}
