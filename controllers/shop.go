package controllers

import (
	"GoShopping/database"
	"GoShopping/models"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func AddShop(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var shopData models.Shop
	err := json.Unmarshal([]byte(data["id"]), &shopData)

	if err != nil {
		return err
	}
	// {
	// 	"id": [1,2,3],
	// 	"quantity":[2,2,3]
	// 	"ownby":"123"
	// }

	// shopData := models.Shop{
	// 	ProductId: data["id"],
	// 	Quantity:  data["quantity"],
	// 	OwnBy:     data["ownby"],
	// }

	database.DB.Create(&data)

	// for i := 0; i < len(data["id"]); i++ {
	// 	user := models.Shop{
	// 		ProductId: data["id"][i],
	// 		Quantity: data["quantity"][i],
	// 		OwnBy: data["ownby"],
	// 	}

	// 	database.DB.Create(&user)
	// }

	return c.JSON(fiber.Map{
		"message": "Successfully",
	})
}
