package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func PatchTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")
	todoID, err := primitive.ObjectIDFromHex(idParam)

	if err != nil {
		return c.Status(400).SendString("Invalid ID format")
	}

	// Parse request body to get fields to update
	type UpdateFields struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		IsCompleted bool   `json:"iscompleted,omitempty"`
	}
	updateFields := new(UpdateFields)
	if err := c.BodyParser(updateFields); err != nil {
		return c.Status(400).SendString("Error parsing request body: " + err.Error())
	}

	// Prepare the update document based on the provided fields
	updateDoc := bson.D{}
	if updateFields.Title != "" {
		updateDoc = append(updateDoc, bson.E{Key: "title", Value: updateFields.Title})
	}
	if updateFields.Description != "" {
		updateDoc = append(updateDoc, bson.E{Key: "description", Value: updateFields.Description})
	}
	updateDoc = append(updateDoc, bson.E{Key: "iscompleted", Value: updateFields.IsCompleted})
	updateDoc = append(updateDoc, bson.E{Key: "updatedat", Value: time.Now()})

	// Perform the update operation
	query := bson.D{{Key: "_id", Value: todoID}}
	update := bson.D{
		{Key: "$set", Value: updateDoc},
	}
	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := mg.Db.Collection("Todo-1").FindOneAndUpdate(c.Context(), query, update, options)

	// Handle errors and return response
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString("Error updating Todo: " + result.Err().Error())
	}

	updatedTodo := new(Todo)
	if err := result.Decode(updatedTodo); err != nil {
		return c.Status(500).SendString("Error decoding updated Todo: " + err.Error())
	}

	updatedTodo.ID = idParam
	return c.Status(200).JSON(updatedTodo)
}
