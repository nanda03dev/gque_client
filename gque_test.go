package gqueclient

import (
	"fmt"
	"testing"
	"time"
)

const (
	GQUE_URI = "localhost:5456"
)

func TestGque(t *testing.T) {
	var AppName = "gque-gque"

	GqueClient := Connect(GQUE_URI, AppName)

	startQueue1(GqueClient)
	startQueue2(GqueClient)

	time.Sleep(2 * time.Second)

	someVar := 1
	index := 1

	pushMessage := QueueMessageType{
		Name: "queue1",
		Data: MessageType{"Id": "Nanda03", "index": 1},
	}

	for range time.Tick(3 * time.Second) {
		if someVar == 1 {
			pushMessage.Name = "queue1"
			someVar = 2
		} else {
			pushMessage.Name = "queue2"
			someVar = 1
		}
		fmt.Printf("\nMessage push to %v ", pushMessage)

		GqueClient.PushMessage(pushMessage)
		index += 1
		pushMessage.Data["index"] = index
	}
}

func startQueue1(GqueClient *Client) {
	queue1 := Queue{
		Name: "queue1",
		Time: 600,
	}

	queue1CreateResult, queue1CreateErr := GqueClient.CreateQueue(queue1)
	fmt.Printf("\n queue1CreateResult : %v \n queue1CreateErr: %v ", queue1CreateResult, queue1CreateErr)

	receiveChan := make(chan MessageType)
	go consumerTest(queue1.Name, receiveChan)

	consumer1Request := ConsumerRequestType{
		QueueName: queue1.Name,
	}
	consumer1RequestErr := GqueClient.Consume(consumer1Request, receiveChan)
	if consumer1RequestErr != nil {
		fmt.Printf("\nconsumer1RequestErr : %v ", consumer1RequestErr)
	}

}
func startQueue2(GqueClient *Client) {

	queue2 := Queue{
		Name: "queue2",
		Time: 600,
	}

	queue2CreateResult, queue2CreateErr := GqueClient.CreateQueue(queue2)
	fmt.Printf("\n queue2CreateResult : %v \n queue2CreateErr: %v ", queue2CreateResult, queue2CreateErr)

	receiveChan2 := make(chan MessageType)

	go consumerTest(queue2.Name, receiveChan2)

	consumer2Request := ConsumerRequestType{
		QueueName: queue2.Name,
	}

	consumer2RequestErr := GqueClient.Consume(consumer2Request, receiveChan2)

	if consumer2RequestErr != nil {
		fmt.Printf("\nconsumer2RequestErr: %v ", consumer2RequestErr)
	}
}

func consumerTest(queueName string, receiveChan chan MessageType) {
	fmt.Printf("\nQueueName : %v consumer started ", queueName)
	for {
		msg, ok := <-receiveChan
		if !ok {
			break
		}
		fmt.Printf("\nQueueName: %v ---- Received Message: %v", queueName, msg)
	}
}
