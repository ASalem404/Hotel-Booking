package api

import (
	"context"

	"hotelSys/db"
	"hotelSys/types"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *userHandler {
	return &userHandler{userStore: userStore}
}

func (h userHandler) HandleGetUsers(c *fiber.Ctx) error {
	var users []*types.User
	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}
	return c.JSON(users)
}
func (h userHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	context := context.Background()
	user, err := h.userStore.GetUserByID(context, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h userHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.UserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	validationErrors := params.Validate()
	if len(validationErrors) > 0 {
		return c.JSON(validationErrors)
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUSer, err := h.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUSer)
}
