package cache

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"
)

//Redis redis cache
type Redis struct {
	conn *redis.Pool
}

//RedisOpts redis conn property
type RedisOpts struct {
	Host        string `yml:"host" json:"host"`
	Password    string `yml:"password" json:"password"`
	Database    int    `yml:"database" json:"database"`
	MaxIdle     int    `yml:"max_idle" json:"max_idle"`
	MaxActive   int    `yml:"max_active" json:"max_active"`
	IdleTimeout int    `yml:"idle_timeout" json:"idle_timeout"` //second
}

//NewRedis instance
func NewRedis(opts *RedisOpts) *Redis {
	pool := &redis.Pool{
		MaxActive:   opts.MaxActive,
		MaxIdle:     opts.MaxIdle,
		IdleTimeout: time.Second * time.Duration(opts.IdleTimeout),
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", opts.Host,
				redis.DialDatabase(opts.Database),
				redis.DialPassword(opts.Password),
			)
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := conn.Do("PING")
			return err
		},
	}
	return &Redis{pool}
}

//SetRedisPool set  redis pool
func (r *Redis) SetRedisPool(pool *redis.Pool) {
	r.conn = pool
}

//SetConn set conn
func (r *Redis) SetConn(conn *redis.Pool) {
	r.conn = conn
}

//Get get a value
func (r *Redis) Get(key string) interface{} {
	conn := r.conn.Get()
	defer conn.Close()

	var data []byte
	var err error
	if data, err = redis.Bytes(conn.Do("GET", key)); err != nil {
		return nil
	}
	var reply interface{}
	if err = json.Unmarshal(data, &reply); err != nil {
		return nil
	}

	return reply
}

//Set set a value
func (r *Redis) Set(key string, val interface{}, timeout time.Duration) (err error) {
	conn := r.conn.Get()
	defer conn.Close()

	var data []byte
	if data, err = json.Marshal(val); err != nil {
		return
	}

	_, err = conn.Do("SETEX", key, int64(timeout/time.Second), data)

	return
}

//IsExist is exist null
func (r *Redis) IsExist(key string) bool {
	conn := r.conn.Get()
	defer conn.Close()

	a, _ := conn.Do("EXISTS", key)
	i := a.(int64)
	return i > 0
}

//Delete delete
func (r *Redis) Delete(key string) error {
	conn := r.conn.Get()
	defer conn.Close()

	if _, err := conn.Do("DEL", key); err != nil {
		return err
	}

	return nil
}

