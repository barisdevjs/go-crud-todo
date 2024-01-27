package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


func GetTodo(c *fiber.Ctx) error {
    todoID, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        fmt.Println("Error parsing ID:", err)
        return c.SendStatus(fiber.StatusBadRequest)
    }

    query := bson.D{{Key: "_id", Value: todoID}}

    result := mg.Db.Collection("Todo-1").FindOne(c.Context(), &query)
    if result.Err() != nil {
        if result.Err() == mongo.ErrNoDocuments {
            fmt.Println("No document found")
            return c.SendStatus(fiber.StatusNotFound)
        }
        fmt.Println("Error during FindOne:", result.Err())
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    var todo Todo
    err = result.Decode(&todo)
    if err != nil {
        fmt.Println("Error decoding document:", err)
        return c.SendStatus(fiber.StatusInternalServerError)
    }

    return c.Status(fiber.StatusOK).JSON(todo)
}
