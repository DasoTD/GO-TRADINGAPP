package main

import (
 "fmt"
 "os"
 "strings"
 "github.com/joho/godotenv"
 trades "dasotd/producer/trades"  // <==== add this
)

func main() {
 err := godotenv.Load("../.env")
 if err != nil {
  fmt.Println("Failed to load environment")
 }
 fmt.Println("E work")
 t := os.Getenv("TICKERS")
 topics := strings.Split(t, ",")
 for i,topic := range topics {
	// fmt.Println(topic)
  topics[i] = strings.Trim(strings.Trim(topic,"\\"),"\"") 
 }

 trades.SubScribeAndListen( // <==== add this
  topics,
 )
}