package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Sushant@123"
	dbname   = "progflow"
)

func ConnectToDB() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("DATABASE_URI")))
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	fmt.Println("Connected to database")

	return client
}

func ConnectToPostgres() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("SQL Connect Error: ", err)
	}

	return db
}
