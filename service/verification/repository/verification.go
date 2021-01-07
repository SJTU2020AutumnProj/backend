package repository

import (
	"context"
	"log"

	"github.com/garyburd/redigo/redis"
	redigo "github.com/garyburd/redigo/redis"
)

/*
VerificationRepository interface
*/
type VerificationRepository interface {
	Set(ctx context.Context, key string, value string, expireTime string) error
	Get(ctx context.Context, key string) (string, error)
}

/*
VerificationRepositoryImpl implementation
*/
type VerificationRepositoryImpl struct {
	CONN redigo.Conn
}

/*
Set set (key, value) pair
*/
func (repo *VerificationRepositoryImpl) Set(ctx context.Context, key string, value string, expireTime string) error {
	_, err := repo.CONN.Do("SET", key, value, "EX", expireTime)
	if nil != err {
		log.Println("VerificationRepository SET key", key, "value", value, "error ", err)
		return err
	}
	return nil
}

/*
Get get value according to key
*/
func (repo *VerificationRepositoryImpl) Get(ctx context.Context, key string) (string, error) {
	result, err := redis.String(repo.CONN.Do("GET", key))
	if err != nil {
		log.Println("VerificationRepository GET key", key, "error", err)
		return "", err
	}
	return result, nil
}
