package publisher

import (
	"context"
	"fmt"

	"github.com/yadavkl/redispubsub/utils"
)

type Publisher struct {
	RedisClient  *utils.RedisClient
	RedisChannel string
	PublisherId  string
}

func NewPublisher(redisClient *utils.RedisClient, redisChannel string, publisherId string) *Publisher {
	return &Publisher{
		RedisClient:  redisClient,
		RedisChannel: redisChannel,
		PublisherId:  publisherId,
	}
}

func (p *Publisher) SendMessage(ctx context.Context, msg interface{}) error {
	err := p.RedisClient.Client.Publish(ctx, p.RedisChannel, msg).Err()
	if err != nil {
		fmt.Println("publisher:", p.PublisherId, "failed to publish message, error:", err)
	}
	return nil

}
