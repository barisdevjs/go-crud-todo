package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	todoID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString("Error parsing request body: " + err.Error())
	}

	query := bson.D{{Key: "_id", Value: todoID}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "title", Value: todo.Title},
				{Key: "description", Value: todo.Description},
				{Key: "iscompleted", Value: todo.IsCompleted},
				{Key: "createdat", Value: todo.CreatedAt},
				{Key: "updatedat", Value: time.Now()},
			},
		},
	}

	// Set the ReturnDocument option to mongo.After to get the updated document
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := mg.Db.Collection("Todo-1").FindOneAndUpdate(c.Context(), query, update, options)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString("Error updating Todo: " + result.Err().Error())
	}

	if err := result.Decode(todo); err != nil {
		return c.Status(500).SendString("Error decoding updated Todo: " + err.Error())
	}

	todo.ID = idParam
	return c.Status(200).JSON(todo)
}
