package mq

import "context"

type Costumer interface {
	Close() error
	Subscribe(ctx context.Context, groupName string, topic string, handlerFunc ...interface{}) error
}
