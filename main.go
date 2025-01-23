package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/yadavkl/redispubsub/publisher"
	"github.com/yadavkl/redispubsub/subscriber"
	"github.com/yadavkl/redispubsub/utils"
)

func main() {
	//getting new redis client
	redis := utils.NewClient("127.0.0.1:6379", "")
	done := make(chan struct{})
	//Create publisher
	pub := publisher.NewPublisher(redis, "temp-channel-redis", "temp-pub-id")
	pub1 := publisher.NewPublisher(redis, "temp1-channel-redis", "temp-pub1-id")
	sub := subscriber.NewSubscriber(redis, []string{"temp-channel-redis", "temp1-channel-redis"}, "temp-sub-id")
	sub.Suscribe(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		msgch, errch := sub.ReceiveMassage(context.Background())
		for {
			select {
			case msg := <-msgch:
				fmt.Println("message received from channel:", msg.Channel, "payload:", msg.Payload)
			case err := <-errch:
				fmt.Println("not able to receeive maessage", err)
				return
			case <-done:
				return
			}
		}
	}(&wg)
	messages := []string{
		"hi",
		"how",
		"are",
		"you",
		"there",
	}
	for _, msg := range messages {
		pub.SendMessage(context.Background(), msg)
	}
	for _, msg := range messages {
		pub1.SendMessage(context.Background(), msg)
	}
	done <- struct{}{}
	wg.Wait()
}
