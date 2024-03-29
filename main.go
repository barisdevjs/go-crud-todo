package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance
var USER_NAME string
var DB_PASS string

const dbName = "Todo"

type Todo struct {
	ID          string     `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsCompleted bool       `json:"is_completed"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// bson means that we see the at bson format at MongoDb
// but it will come like a json format

func Connect() error {
	currentWorkDirectory, _ := os.Getwd()
	fmt.Println("Current working directory:", currentWorkDirectory)

	err := godotenv.Load(currentWorkDirectory + "/.env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		log.Fatal("Error loading .env file 11")
	}

	USER_NAME = os.Getenv("USER_NAME")
	DB_PASS = os.Getenv("DB_PASS")

	escapedUser := url.QueryEscape(USER_NAME)
	escapedPass := url.QueryEscape(DB_PASS)
	// mongoURI := "mongodb+srv://" + escapedUser + ":" + escapedPass + "@cluster0.smzb3os.mongodb.net/" + dbName + "?retryWrites=true&w=majority"
	mongoURI := "mongodb+srv://" + escapedUser + ":" + escapedPass + "@cluster0.smzb3os.mongodb.net/?retryWrites=true&w=majority&directConnection=false"
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	db := client.Database(dbName)
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	fmt.Println(mongoURI)
	return nil
}

// fiber ~~ like express.js
func main() {

	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()

	app.Get("/getAllTodos", GetTodos)
	app.Get("/getTodo/:id", GetTodo)
	app.Post("/createTodo", CreateTodo)
	app.Delete("/deleteTodo/:id", DeleteTodo)
	app.Delete("/deleteTodos", DeleteTodos)
	app.Put("/updateTodo/:id", UpdateTodo)
	app.Patch("/updatePartialTodo/:id", PatchTodo)
	app.Options("/options", GetOptions)
	

	log.Fatal(app.Listen(":8085"))

}
