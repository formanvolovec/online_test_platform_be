package main

import (
	"context"
	"fmt"
	"log"
	"otfch_be/db_connection"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Could not connect to Redis: %v", err))
	}

	return rdb
}

func main() {
	//Initializing the Application
	app := fiber.New()
	//Initializing the Redis
	rdb := initRedis()

	app.Get("/", func(c *fiber.Ctx) error {
		err := rdb.Set(ctx, "key", "success", 0).Err()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.SendString(val)
	})

	app.Listen(":3000")

	//Initializing the database
	pool, err := db_connection.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer pool.Close()
}
