package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RequestBody struct {
	IDs []string `json:"ids"`
}

type DeleteResponse struct {
	DeletedIDs []string `json:"deleted_ids"`
	Statuses   map[string]bool `json:"statuses"`
}

func DeleteTodos(c *fiber.Ctx) error {
	var reqBody RequestBody
	if err := c.BodyParser(&reqBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Failed to parse request body"})
	}

	// Get the IDs from the request body
	idArray := reqBody.IDs

	// Create a slice to hold the ObjectIDs
	var objectIDs []primitive.ObjectID

	// Convert each string ID to primitive.ObjectID and add it to the slice
	for _, id := range idArray {
		employeeID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid ID format"})
		}
		objectIDs = append(objectIDs, employeeID)
	}

	// Get the MongoDB collection
	collection := mg.Db.Collection("Todo-1")

	// Create the filter with $in operator to match multiple employee IDs
	filter := bson.M{"_id": bson.M{"$in": objectIDs}}

	// Perform the DeleteMany operation
	result, err := collection.DeleteMany(c.Context(), filter)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete records", "details": err.Error()})
	}

	// Prepare the response
	response := DeleteResponse{
		DeletedIDs: make([]string, 0),
		Statuses:   make(map[string]bool),
	}

	// Populate the response with deleted IDs and their statuses
	for _, id := range objectIDs {
		response.DeletedIDs = append(response.DeletedIDs, id.Hex())
	}

	// Populate the statuses
	if result.DeletedCount > 0 {
		for _, id := range response.DeletedIDs {
			response.Statuses[id] = true
		}
	}

	// Return the response
	return c.Status(200).JSON(response)
}
