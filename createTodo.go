package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateTodo(c *fiber.Ctx) error {
	collection := mg.Db.Collection("Todo-1")

	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	todo.ID = ""

	fmt.Printf("Received Todo: %+v\n", todo)

	insertionResult, err := collection.InsertOne(c.Context(), todo)

	if err != nil {
		fmt.Printf("Error inserting Todo: %s\n", err.Error())
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
	createdRecord := collection.FindOne(c.Context(), filter)

	createdTodo := &Todo{}
	createdRecord.Decode(createdTodo)

	return c.Status(201).JSON(createdTodo)
}
