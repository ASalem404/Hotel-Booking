package main

import (
	"context"
	"hotelSys/api"
	"hotelSys/db"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBURI = "mongodb://localhost:27017"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// coll := client.Database(dbName).Collection(dbColl)

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	app := fiber.New(config)
	apiV1Grp := app.Group("/api/v1")
	apiV1Grp.Get("/user", userHandler.HandleGetUsers)
	apiV1Grp.Get("/user/:id", userHandler.HandleGetUser)
	apiV1Grp.Post("/user", userHandler.HandlePostUser)
	app.Listen(":8080")
}
