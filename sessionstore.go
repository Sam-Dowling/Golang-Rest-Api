package main

import (
	"gopkg.in/redis.v3"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func addSession(token string) {
	err := client.Set(token, "", time.Hour*24).Err()
	checkErr(err)
}

func getSession(token string) bool {
	val, err := client.Exists(token).Result()
	checkErr(err)
	return val
}
