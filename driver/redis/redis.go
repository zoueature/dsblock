// Package dsblock
// Author: Zoueature
// Email: zoueature@gmail.com
// -------------------------------

package redis

import (
	"github.com/go-redis/redis"
	"github.com/zoueature/dsblock"
	"time"
)

// NewSingleLocker 单实例的分布式锁
func NewSingleLocker(addr, password string, db int) dsblock.Locker {
	conn := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redisLock{conn: conn}
}

type redisLock struct {
	conn *redis.Client
}

func (r *redisLock) Lock(key string, autoUnlockTime time.Duration) error {
	return r.conn.SetNX(key, "OK", autoUnlockTime).Err()
}

func (r *redisLock) UnLock(key string) error {
	return r.conn.Del(key).Err()
}

func (r *redisLock) Close() {
	_ = r.conn.Close()
}
