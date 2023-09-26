package main

import (
	"context"
	"fmt"
	"log"
	"time"

	redisDriver "github.com/redis/go-redis/v9"
)

func newRedisClient(host string, password string) *redisDriver.Client {
	client := redisDriver.NewClient(&redisDriver.Options{
		Addr:     host,
		Password: password,
		DB:       0,
	})
	return client
}

func setData(rdc *redisDriver.Client, key string, data string, ttl time.Duration) error {
	dataSet := rdc.Set(context.Background(), key, data, ttl)
	return dataSet.Err()
}

func getData(rdc *redisDriver.Client, key string) (string, error) {
	dataGet := rdc.Get(context.Background(), key)
	if dataGet.Err() != nil {
		fmt.Printf("data tidak ditemukan : %v", dataGet.Err())
		return "", dataGet.Err()
	}
	resp, err := dataGet.Result()
	return resp, err
}

func main() {
	var host = "localhost:63790"
	var redisPassword = "redis"
	rdc := newRedisClient(host, redisPassword)

	key := "sample-test"
	data := "halo ini test redis pertama sama"
	ttl := time.Duration(60) * time.Second

	//store data
	errSetData := setData(rdc, key, data, ttl)
	if errSetData != nil {
		fmt.Printf("set data error: %v", errSetData)
		return
	}
	log.Println("set data success")

	//get data
	getdata, err := getData(rdc, key)
	if err != nil {
		fmt.Println("get data error : %v", err)
		return
	}
	log.Println("get data success result:", getdata)
}
