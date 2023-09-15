package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	client.Set("language", "Go", 0)
	language := client.Get("language")
	year := client.Get("year")

	fmt.Println(language.Val())
	fmt.Println(year.Val())
}
