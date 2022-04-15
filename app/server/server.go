package server

import "os"

func GetConnectionString() string {
	dbName := os.Getenv("DB_DATABASE_NAME")
	return "root:@tcp(127.0.0.1:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
