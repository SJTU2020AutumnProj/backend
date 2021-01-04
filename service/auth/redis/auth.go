package redis

import (
	"context"
	"log"

	"github.com/garyburd/redigo/redis"
	redigo "github.com/garyburd/redigo/redis"
)

// AuthRedis interface
type AuthRedis interface {
	Set(ctx context.Context, key string, value string, expireTime string) error
	Get(ctx context.Context, key string) (string, error)
}

// AuthRedisImpl implementation struct
type AuthRedisImpl struct {
	CONN redigo.Conn
}

/*
Set set (key, value) pair
*/
func (repo *AuthRedisImpl) Set(ctx context.Context, key string, value string, expireTime string) error {
	_, err := repo.CONN.Do("SET", key, value, "EX", expireTime)
	if nil != err {
		log.Println("AuthRedis Set key", key, "value", value, "error ", err)
		return err
	}
	return nil
}

/*
Get get value according to key
*/
func (repo *AuthRedisImpl) Get(ctx context.Context, key string) (string, error) {
	result, err := redis.String(repo.CONN.Do("GET", key))
	if err != nil {
		log.Println("AuthRedis Get key", key, "error", err)
		return "", err
	}
	return result, nil
}



