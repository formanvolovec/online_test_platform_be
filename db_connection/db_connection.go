package db_connection

import (
  "context"
  "log"
  "os"
  "github.com/joho/godotenv"
  "github.com/jackc/pgx/v4/pgxpool"
)

func InitDB() (*pgxpool.Pool, error) {
  // Загружаем переменные окружения из .env файла
  err := godotenv.Load()
  if err != nil {
    return nil, err
  }

  // Получаем значение переменной окружения DATABASE_URL
  databaseURL := os.Getenv("DATABASE_URL")

  poolConfig, err := pgxpool.ParseConfig(databaseURL)
  if err != nil {
    return nil, err
  }

  pool, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
  if err != nil {
    return nil, err
  }

  log.Println("Successfully connected to the database")
  return pool, nil
}
