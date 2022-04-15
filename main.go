package main

import (
	"randi_firmansyah/app/helper/helper"
	"randi_firmansyah/app/server"

	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load("params/.env")
	helper.CheckEnv(err)

	// running server
	server.Execute()
}
