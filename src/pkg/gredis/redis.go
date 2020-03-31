package gredis

import (
	"time"
	"goms/pkg/setting"
	"goms/pkg/logging"
	"github.com/gomodule/redigo/redis"
)

var RedisCon *redis.Pool

func SetUp() error {
	RedisCon = &redis.Pool{
		MaxIdle: setting.RedisSetting.MaxIdle,
		MaxActive: setting.RedisSetting.MaxActive,
		IdleTimeout: setting.RedisSetting.IdleTimeout,
		Wait: true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", setting.RedisSetting.Host)
			if err != nil {
				logging.Error(err)
				return nil, err
			}
			if setting.RedisSetting.Password != "" {
				if _, err := c.Do("AUTH", setting.RedisSetting.Password); err != nil {
					c.Close()
					logging.Error(err)
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				logging.Error(err)
			}
			return err
		},
	}
	return nil
}

// Set a key/value
func Set(key, value string, time int) error {
	conn := RedisCon.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	if time > 0 {
		_, err = conn.Do("EXPIRE", key, time)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get get a key
func Get(key string) ([]byte, error) {
	conn := RedisCon.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}