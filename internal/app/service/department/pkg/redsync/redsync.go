package redsync

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"resume-resolving/internal/app/service/department/config"
	"sync"
	"time"
)

const cancelTime = 5 * time.Second

type RedSync struct {
	RedisPool *redsync.Redsync
	LockPool  *sync.Pool
	config    *config.Config
}

func New(config2 *config.Config) *RedSync {
	return &RedSync{
		config: config2,
	}
}

func (r *RedSync) Init() error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", r.config.ConfigInNacos.Redis.Host, r.config.ConfigInNacos.Redis.Port),
		Password: r.config.ConfigInNacos.Redis.Password,
		DB:       r.config.ConfigInNacos.Redis.Db,
		PoolSize: r.config.ConfigInNacos.Redis.PoolSize,
	})
	ctx, cancel := context.WithTimeout(context.Background(), cancelTime)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}

	pool := goredis.NewPool(rdb)
	r.RedisPool = redsync.New(pool)

	r.LockPool = &sync.Pool{
		New: func() interface{} {
			return new(redsync.Mutex)
		},
	}

	return nil
}

func (r *RedSync) Close() error {
	return nil
}
