package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"log"
	"net/http"
	"os"
	"database/sql"

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

	// create a database object, which manages a pool of
	// network connections to the database server
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("error opening database: %v\n", err)
		os.Exit(1)
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging database: %v\n", err)
	} else {
		fmt.Printf("Successfully connected!\n")
	}

	context := &handlers.HandlerContext {

	}

	mux := http.NewServeMux()

	// Handler functions...

	fmt.Printf("listening on %s... \n", addr)
	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, mux))
}
