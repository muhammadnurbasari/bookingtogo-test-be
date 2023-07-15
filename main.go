package main

import (
	"fmt"
	"os"

	"bookingtogo-test-be/connectDB"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Read file env. . .")
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Error().Msg(err.Error())
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	conn, errConn := connectDB.ConnMySQL(dbHost, dbPort, dbUser, dbPass, dbName)

	if errConn != nil {
		log.Error().Msg(errConn.Error())
	}
	fmt.Println(conn)
}
