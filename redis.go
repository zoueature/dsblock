// Package dsblock
// Author: Zoueature           
// Email: zoueature@gmail.com  
// -------------------------------

package dsblock

import "github.com/go-redis/redis"

func RedisLocker() DsbLock {
    conn := redis.NewClient(&redis.Options{
        Network:            "",
        Addr:               "",
        Dialer:             nil,
        OnConnect:          nil,
        Password:           "",
        DB:                 0,
        MaxRetries:         0,
        MinRetryBackoff:    0,
        MaxRetryBackoff:    0,
        DialTimeout:        0,
        ReadTimeout:        0,
        WriteTimeout:       0,
        PoolSize:           0,
        MinIdleConns:       0,
        MaxConnAge:         0,
        PoolTimeout:        0,
        IdleTimeout:        0,
        IdleCheckFrequency: 0,
        TLSConfig:          nil,
    })
    return &redisLock{conn: conn}
}

type redisLock struct {
    conn *redis.Client
}

func (r *redisLock) Lock() error {
    panic("implement me")
}

func (r *redisLock) UnLock() error {
    panic("implement me")
}

func (r *redisLock) Close() {
    panic("implement me")
}

