package mq

import (
	"context"
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"resume-resolving/internal/app/service/department/config"
	"resume-resolving/internal/pkg/mq"
)

const (
	errTypeAssertFailed = "assert type failed"
)

type RocketmqHandlerFunc func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)

type RocketmqProducer struct {
	rocketmqProducer rocketmq.Producer
	config           *config.Config
}

func (r *RocketmqProducer) Open(groupName string) error {
	p, err := rocketmq.NewProducer(
		producer.WithNameServer([]string{fmt.Sprintf("%s:%d",
			r.config.ConfigInNacos.Rocketmq.Host,
			r.config.ConfigInNacos.Rocketmq.Port)}),
		producer.WithGroupName(groupName),
		producer.WithRetry(r.config.ConfigInNacos.Rocketmq.RetryCount),
	)
	r.rocketmqProducer = p
	if err != nil {
		return err
	}
	err = r.rocketmqProducer.Start()
	if err != nil {
		return err
	}
	return nil
}

func (r *RocketmqProducer) Close() error {
	return r.rocketmqProducer.Shutdown()
}

func (r *RocketmqProducer) Send(ctx context.Context, topic string, msg ...[]byte) (err error) {
	message := &primitive.Message{
		Topic: topic,
		Body:  msg[0],
	}
	_, err = r.rocketmqProducer.SendSync(context.Background(), message)
	return
}

type RocketmqCostumer struct {
	rocketmqConsumer rocketmq.PushConsumer
	config           *config.Config
}

func (r *RocketmqCostumer) Subscribe(ctx context.Context, groupName, topic string, handlerFunc ...interface{}) error {
	handler, ok := handlerFunc[0].(RocketmqHandlerFunc)
	if !ok {
		return errors.New(errTypeAssertFailed)
	}
	c, err := rocketmq.NewPushConsumer(
		// 指定 Group 可以实现消费者负载均衡进行消费，并且保证他们的Topic+Tag要一样。
		// 如果同一个 GroupID 下的不同消费者实例，订阅了不同的 Topic+Tag 将导致在对Topic 的消费队列进行负载均衡的时候产生不正确的结果，最终导致消息丢失。(官方源码设计)
		consumer.WithGroupName(groupName),
		consumer.WithNameServer([]string{fmt.Sprintf("%s:%d",
			r.config.ConfigInNacos.Rocketmq.Host,
			r.config.ConfigInNacos.Rocketmq.Port)}),
	)
	r.rocketmqConsumer = c
	err = r.rocketmqConsumer.Subscribe(topic, consumer.MessageSelector{}, handler)
	if err != nil {
		return err
	}
	err = r.rocketmqConsumer.Start()
	if err != nil {
		return err
	}
	return nil
}

func (r *RocketmqCostumer) Close() error {
	return r.rocketmqConsumer.Shutdown()
}

func NewRocketmqProducer(config *config.Config) mq.Producer {
	return &RocketmqProducer{
		config: config,
	}
}

func NewRocketmqCostumer(config *config.Config) mq.Costumer {
	return &RocketmqCostumer{
		config: config,
	}
}
