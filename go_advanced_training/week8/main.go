package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var valueLength = []int{10, 20, 50, 100, 200, 1000, 5000, 10000}

const COUNT = 200000

func main() {
	rand.Seed(time.Now().UnixNano())
	keys := make([]string, COUNT)
	for i := range keys {
		keys[i] = RandStringRunes(15)
	}
	client := redis.NewClient(&redis.Options{Addr: "redis-server:6379"})
	if err := client.FlushAll(context.Background()).Err(); err != nil {
		panic(err)
	}
	log.Println("| size  | before | after | actual | per | about |")
	log.Println("| :---: | :----: | :---: | :----: | :-: | :---: |")
	for _, size := range valueLength {
		before, err := GetMemory(client)
		if err != nil {
			panic(err)
		}
		for _, key := range keys {
			if err := client.Set(context.Background(), key, RandStringRunes(size), 0).Err(); err != nil {
				panic(err)
			}
		}
		after, err := GetMemory(client)
		if err != nil {
			panic(err)
		}
		log.Printf("|%v|%v|%v|%v|%v|%v|", size, before, after, after-before, float64(after-before)/float64(COUNT), math.Round(float64(after-before)/float64(COUNT)))
		if err := client.FlushAll(context.Background()).Err(); err != nil {
			panic(err)
		}
	}

}

func GetMemory(client *redis.Client) (int64, error) {
	c, err := client.Info(context.Background(), "memory").Result()
	if err != nil {
		return 0, err
	}
	memoryDataset := strings.Split(c, "\r\n")[10]
	memoryStr := strings.Split(memoryDataset, ":")[1]
	return strconv.ParseInt(memoryStr, 10, 64)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
