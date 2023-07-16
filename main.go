package main

import (
	"net/http"
	"os"

	"bookingtogo-test-be/connectDB"
	"bookingtogo-test-be/modules/customers/customerHandler"
	"bookingtogo-test-be/modules/customers/customerRepository"
	"bookingtogo-test-be/modules/customers/customerUsecase"

	"github.com/gorilla/mux"
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
	port := os.Getenv("PORT")

	conn, errConn := connectDB.ConnMySQL(dbHost, dbPort, dbUser, dbPass, dbName)

	if errConn != nil {
		log.Error().Msg(errConn.Error())
	}

	r := mux.NewRouter()

	// initialize module
	customerRepo := customerRepository.NewCustomerRepository(conn)
	customerUC := customerUsecase.NewCustomerUsecase(customerRepo)
	customerHandler.NewCustomerHandler(r, customerUC)

	http.Handle("/", r)
	log.Info().Msg("Server run on port " + port)
	http.ListenAndServe(":8080", nil)
}
