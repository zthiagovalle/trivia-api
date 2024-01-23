package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zthiagovalle/trivia-api/database"
	"github.com/zthiagovalle/trivia-api/models"
)

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}

	database.DB.Db.Find(&facts)

	return c.Status(http.StatusOK).JSON(facts)
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(http.StatusOK).JSON(fact)
}
