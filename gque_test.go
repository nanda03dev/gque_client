package gqueclient

import (
	"fmt"
	"testing"
	"time"
)

const (
	GQUE_URI = "localhost:5456"
)

var Queue1 = Queue{
	Name: "queue1",
	Time: 600,
}

var Queue2 = Queue{
	Name: "queue2",
	Time: 600,
}

var Broadcast1 = Broadcast{
	Name:       "all-queue",
	QueueNames: []string{Queue1.Name, Queue2.Name},
}
var (
	QUEUE_TIME_INTERVAL     = 100 * time.Millisecond
	BROADCAST_TIME_INTERVAL = 300 * time.Millisecond
)

func TestGque(t *testing.T) {
	var AppName = "gque-gque"

	GqueClient := Connect(GQUE_URI, AppName)

	StartQueue1(GqueClient)
	StartQueue2(GqueClient)
	go StartBroadcast(GqueClient)

	time.Sleep(2 * time.Second)

	someVar := 1
	index := 1

	pushMessage := QueueMessageType{
		Name: "queue1",
		Data: MessageType{"MessageFrom": "Queue", "index": 1},
	}

	for range time.Tick(QUEUE_TIME_INTERVAL) {
		if someVar == 1 {
			pushMessage.Name = Queue1.Name
			someVar = 2
		} else {
			pushMessage.Name = Queue2.Name
			someVar = 1
		}
		fmt.Printf("\nMessage push to %v ", pushMessage)
		pushMessage.Data["index"] = index

		GqueClient.PushMessage(pushMessage)
		index += 1
	}
}

func StartQueue1(GqueClient *Client) {

	queue1CreateResult, queue1CreateErr := GqueClient.CreateQueue(Queue1)
	fmt.Printf("\n queue1CreateResult : %v \n queue1CreateErr: %v ", queue1CreateResult, queue1CreateErr)

	receiveChan := make(chan MessageType, 10000)
	go ConsumerTest(Queue1.Name, receiveChan)

	consumer1Request := ConsumerRequestType{
		QueueName: Queue1.Name,
	}
	consumer1RequestErr := GqueClient.Consume(consumer1Request, receiveChan)
	if consumer1RequestErr != nil {
		fmt.Printf("\nconsumer1RequestErr : %v ", consumer1RequestErr)
	}

}

func StartQueue2(GqueClient *Client) {
	queue2CreateResult, queue2CreateErr := GqueClient.CreateQueue(Queue2)
	fmt.Printf("\n queue2CreateResult : %v \n queue2CreateErr: %v ", queue2CreateResult, queue2CreateErr)

	receiveChan2 := make(chan MessageType, 10000)

	go ConsumerTest(Queue2.Name, receiveChan2)

	consumer2Request := ConsumerRequestType{
		QueueName: Queue2.Name,
	}

	consumer2RequestErr := GqueClient.Consume(consumer2Request, receiveChan2)

	if consumer2RequestErr != nil {
		fmt.Printf("\nconsumer2RequestErr: %v ", consumer2RequestErr)
	}
}

func StartBroadcast(GqueClient *Client) {

	broadcastCreateResult, broadcastCreateErr := GqueClient.CreateBroadcast(Broadcast1)
	fmt.Printf("\n broadcastCreateResult : %v \n broadcastCreateErr: %v ", broadcastCreateResult, broadcastCreateErr)

	broadcastMessage := BroadcastMessageType{
		Name: Broadcast1.Name,
		Data: MessageType{"MessageFrom": "Brodcast", "index": 1},
	}
	time.Sleep(2 * time.Second)
	index := 1
	for range time.Tick(BROADCAST_TIME_INTERVAL) {
		broadcastMessage.Data["index"] = index
		GqueClient.BroadcastMessage(broadcastMessage)

		index += 1

	}
}

func ConsumerTest(queueName string, receiveChan chan MessageType) {
	fmt.Printf("\nQueueName : %v consumer started ", queueName)
	for {
		msg, ok := <-receiveChan
		if !ok {
			break
		}
		fmt.Printf("\nQueueName: %v ---- Received Message: %v", queueName, msg)
	}
}
