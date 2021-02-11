package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DBConnetion = ""
	APIPort     = 0
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// load port api
	APIPort = 5000
	if num, err := strconv.Atoi(os.Getenv("API_PORT")); err == nil {
		APIPort = num
	}

	// load url connection bd
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	schema := os.Getenv("DB_SCHEMA")

	DBConnetion = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, schema)
}
