package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	if len(tlsKeyPath) < 0 || len(tlsCertPath) < 0 {
		log.Fatal("No environment variable found for either TLSKey or TLSCERT")
	}

	sessionKey := os.Getenv("SESSIONKEY")
	reddisAddr := os.Getenv("REDISADDR")
	dsn := os.Getenv("DSN")

	// Starts a new redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:	reddisAddr,
		Password: "",
		DB: 0,
	})

	db, err := sql.Open("mysql", dsn)
	if err != {
		fmt.Println("There was an error opening the database.")
	} else {
		fmt.Println("Database successfully opened!")
	}


	mux := http.NewServeMux()

	fmt.Printf("listening on %s... \n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}
