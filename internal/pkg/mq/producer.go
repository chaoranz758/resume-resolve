package mq

import (
	"context"
)

type Producer interface {
	Open(groupName string) error
	Close() error
	Send(ctx context.Context, topic string, msg ...[]byte) error
}
