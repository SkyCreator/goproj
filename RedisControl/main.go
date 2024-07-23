package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello, RedisControl!")
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	redisSec := cfg.Section("redis")
	strSec := cfg.Section("FuzzyStr")
	ip := redisSec.Key("IP").String()
	port := redisSec.Key("PORT").String()
	pwd := redisSec.Key("PASSWORD").String()
	db, err := redisSec.Key("DB").Int()
	if err != nil {
		fmt.Printf("Fail to read key [DB]: %v", err)
		os.Exit(1)
	}
	ctx := context.Background()
	cli := redis.NewClient(&redis.Options{Addr: ip + ":" + port, Password: pwd, DB: db})
	mapKey := strSec.KeysHash()
	for _, val := range mapKey {
		iter := cli.Scan(ctx, 0, "*"+val+"*", 0).Iterator()
		for iter.Next(ctx) {
			key := iter.Val()
			if err := cli.Del(ctx, key).Err(); err != nil {
				log.Printf("Failed to delete key %s: %v", key, err)
			} else {
				fmt.Printf("Deleted key: %s\n", key)
			}
		}
		if err = iter.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
