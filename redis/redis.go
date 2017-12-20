package main

import (
	"fmt"
	"log"
	"time"

	goredis "gopkg.in/redis.v3"
)

type NewRedis struct {
	conn *goredis.Client
}

func NewConnect(addr string, password string, db int64) (*NewRedis, error) {
	rs := &NewRedis{}
	err := rs.Connect(addr, password, db)
	if err != nil {
		log.Fatal(err)
	}
	return rs, nil

}

func (this *NewRedis) Connect(addr string, password string, db int64) error {
	client := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
		return err
	}
	if pong != "PONG" {
		log.Fatal("connection redis return is not Pong.")
		return err
	}
	this.conn = client
	return nil
}

func (this *NewRedis) Set(k string, v string, expire time.Duration) error {
	if err := this.conn.Set(k, v, expire).Err(); err != nil {
		return err
	}
	return nil
}

func (this *NewRedis) Get(k string) (string, bool, error) {
	if val, err := this.conn.Get(k).Result(); err != nil {
		return "", false, err
	} else if err == goredis.Nil {
		return "", false, nil
	} else {
		return val, true, nil
	}
}

func (this *NewRedis) Del(k string) error {
	if err := this.conn.Del(k).Err(); err != nil {
		return err
	}
	return nil
}

func (this *NewRedis) Lpush(k string, v string) error {
	if err := this.conn.LPush(k, v).Err(); err != nil {
		return err
	}
	return nil
}

func (this *NewRedis) Lpop(k string) error {
	if err := this.conn.LPop(k).Err(); err != nil {
		return err
	}
	return nil

}

func (this *NewRedis) GetKeyLen(k string) (int64, error) {
	if val, err := this.conn.LLen(k).Result(); err != nil {
		return 0, err
	} else {
		return val, nil
	}

}
func (this *NewRedis) Lrange(k string, start, stop int64) ([]string, error) {
	if result, err := this.conn.LRange(k, start, stop).Result(); err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func main() {
	redis, _ := NewConnect("localhost:6379", "", 0)

	/*
		redis.Set("name1", "monkey1", 0)

		v, b, err := redis.Get("namexxxxxxxxx")
		fmt.Println(v)
		fmt.Println(b)
		fmt.Println(err)

		err = redis.Del("name1")
		if err != nil {
			log.Fatal(err)
		}
	*/

	//err := redis.Lpush("arr", "1")
	//fmt.Println(err)

	//keylen, _ := redis.GetKeyLen("arr")
	//fmt.Println(keylen)

	//err := redis.Lpop("arr")
	//fmt.Println(err)

	result, err := redis.Lrange("a1", 1, 100)
	fmt.Println(result)
	fmt.Println(err)
}
