package routes

import (
	"errors"

	"github.com/PrinceNarteh/go-ecommerce-api/database"
	"github.com/PrinceNarteh/go-ecommerce-api/models"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func CreateResponseuser(userModel models.User) User {
	return User{
		ID:        userModel.ID,
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseuser(user)

	return c.Status(201).JSON(responseUser)
}

func GetAllUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseuser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user cannot be found")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	var user models.User
	userId, err := c.ParamsInt("userId")
	if err != nil {
		return c.Status(400).JSON("Please ensure that ID in an integer")
	}

	if err := findUser(userId, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}

	responseUser := CreateResponseuser(user)
	return c.Status(200).JSON(responseUser)
}
