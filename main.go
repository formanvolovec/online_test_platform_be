package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"otfch_be/packages/db_connection"
	"strconv"

	//"otfch_be/packages/s3_connection"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func initRedis() *redis.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	redisDB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Invalid REDIS_DB value: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_SERVER"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       redisDB,
	})

	_, err = rdb.Ping(ctx).Result()
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

	// //Initializing the AWS S3
	// а действительно ли нужно, сложно пиздец я нихуя не понимаю
	// s3_connection.InitS3()

	//Initializing the database
	pool, err := db_connection.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer pool.Close()

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
}
