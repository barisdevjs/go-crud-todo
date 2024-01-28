package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodo(c *fiber.Ctx) error {
	todoID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		fmt.Println("Error parsing todoID:", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	query := bson.D{{Key: "_id", Value: todoID}}

	result, err := mg.Db.Collection("Todo-1").DeleteOne(c.Context(), &query)
	if err != nil {
		fmt.Println("Error deleting document:", err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}


	if result.DeletedCount == 0 {
		fmt.Println("Document not found for deletion")
		return c.SendStatus(fiber.StatusNotFound)
	}

	response := struct {
		ID     string `json:"id"`
		Status bool   `json:"status"`
	}{
		ID:     todoID.Hex(),
		Status: true, // Indicating the deletion was successful
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
