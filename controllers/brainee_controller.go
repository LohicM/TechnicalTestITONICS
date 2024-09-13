package controllers

import (
	"TechnicalTestITONICS/configs"
	"TechnicalTestITONICS/models"
	"TechnicalTestITONICS/responses"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var braineesCollection *mongo.Collection = configs.GetCollection(configs.DB, "brainees")
var validate = validator.New()

func CreateBrainee(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var brainee models.CreateBrainee

	if err := c.BodyParser(&brainee); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Status: http.StatusBadRequest, Error: err.Error()})
	}

	if validationErr := validate.Struct(&brainee); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Status: http.StatusBadRequest, Error: validationErr.Error()})
	}

	result, err := braineesCollection.InsertOne(ctx, brainee)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(responses.BraineeResponse{Status: http.StatusCreated, Data: models.Brainee{
		Id:     result.InsertedID.(primitive.ObjectID),
		Text:   brainee.Text,
		Author: brainee.Author,
		Brand:  brainee.Brand,
	}})
}

func UpdateBrainee(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	braineeId := c.Params("braineeId")
	var brainee models.UpdateBrainee

	objId, _ := primitive.ObjectIDFromHex(braineeId)

	if err := c.BodyParser(&brainee); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Status: http.StatusBadRequest, Error: err.Error()})
	}

	if validationErr := validate.Struct(&brainee); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ErrorResponse{Status: http.StatusBadRequest, Error: validationErr.Error()})
	}

	update := bson.M{}

	if brainee.Text != nil {
		update["text"] = brainee.Text
	}
	if brainee.Author != nil {
		update["author"] = brainee.Author
	}
	if brainee.Brand != nil {
		update["brand"] = brainee.Brand
	}

	result, err := braineesCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
	}

	var updatedBrainee models.Brainee
	if result.MatchedCount == 1 {
		err := braineesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedBrainee)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
		}
	}

	return c.Status(http.StatusOK).JSON(responses.BraineeResponse{Status: http.StatusOK, Data: updatedBrainee})

}

func DeleteBrainee(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	braineeId := c.Params("braineeId")

	objId, _ := primitive.ObjectIDFromHex(braineeId)

	result, err := braineesCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(responses.ErrorResponse{Status: http.StatusNotFound, Error: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(responses.MessageResponse{Status: http.StatusOK, Message: "Brainee deleted."})
}

func GetBraineeById(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	braineeId := c.Params("braineeId")
	var brainee models.Brainee

	objId, _ := primitive.ObjectIDFromHex(braineeId)

	err := braineesCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&brainee)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
	}

	return c.Status(http.StatusOK).JSON(responses.BraineeResponse{Status: http.StatusOK, Data: brainee})
}

func GetBraineesByAuthor(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	braineesAuthor := c.Params("braineesAuthor")
	var brainees []models.Brainee

	results, err := braineesCollection.Find(ctx, bson.M{"author": braineesAuthor})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var brainee models.Brainee
		if err = results.Decode(&brainee); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
		}
		brainees = append(brainees, brainee)
	}
	if brainees == nil {
		return c.Status(http.StatusNotFound).JSON(responses.ErrorResponse{Status: http.StatusNotFound, Error: "Author name not found"})
	}

	return c.Status(http.StatusOK).JSON(responses.BraineesResponse{Status: http.StatusOK, Data: brainees})
}

func GetBraineesByBrand(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	braineesBrand := c.Params("braineesBrand")
	var brainees []models.Brainee

	results, err := braineesCollection.Find(ctx, bson.M{"brand": braineesBrand})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var brainee models.Brainee
		if err = results.Decode(&brainee); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
		}
		brainees = append(brainees, brainee)
	}
	if brainees == nil {
		return c.Status(http.StatusNotFound).JSON(responses.ErrorResponse{Status: http.StatusNotFound, Error: "Brand name not found"})
	}

	return c.Status(http.StatusOK).JSON(responses.BraineesResponse{Status: http.StatusOK, Data: brainees})
}

func GetAllBrainees(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var brainees []models.Brainee

	results, err := braineesCollection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var brainee models.Brainee
		if err = results.Decode(&brainee); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ErrorResponse{Status: http.StatusInternalServerError, Error: err.Error()})
		}
		brainees = append(brainees, brainee)
	}
	return c.Status(http.StatusOK).JSON(responses.BraineesResponse{Status: http.StatusOK, Data: brainees})
}
