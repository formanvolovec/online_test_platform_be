package main

import (
  "fmt"
  "log"
  "net/http"
  "otfch_be/db_connection"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello, World!")
}

func main() {
  pool, err := db_connection.InitDB() 
  if err != nil {
    log.Fatalf("Error initializing database: %v", err)
  }
  defer pool.Close()

  http.HandleFunc("/", helloHandler)
  fmt.Println("Server is running on http://localhost:8080")
  err = http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println("Error starting server:", err)
  }
}
