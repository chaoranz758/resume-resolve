package pkg

import (
	"math/rand"
	"time"
)

func SetRandomExpireTime(random, len int, baseTime time.Duration) []time.Duration {
	r := rand.Intn(random)

	expireTimeList := make([]time.Duration, 0, len)
	for i := 0; i < len; i++ {
		expireTimeList = append(expireTimeList, baseTime+time.Duration(r)*time.Second)
	}
	return expireTimeList
}
