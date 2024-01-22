package handlers

import "github.com/gofiber/fiber/v2"

/*func CreateCredential(c *fiber.Ctx) error {
	var credentials *models.CreateCredentials

	err := c.BodyParser(&credentials)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": err.Error(),
		})
	}

	errors := utils.ValidateStruct(credentials)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	newCredential := models.Credentials{
		Username:     credentials.Username,
		Password:     credentials.Password,
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientSecret,
	}

	db := database.Database.GetDb()

	result := db.Create()
}
*/

type ApiParameters struct {
	Crede
}

func CreateCredential(c *fiber.Ctx) error {
	credential := []{}

	if err := c.QueryParser(&credential); err != nil {
		return err
	}


	return nil
}
