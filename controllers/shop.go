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

	//var shopData models.Shop

	var listProductId []uint
	err1 := json.Unmarshal([]byte(data["id"]), &listProductId)
	if err1 != nil {
		return err1
	}

	var listQuantity []int32
	err2 := json.Unmarshal([]byte(data["quantity"]), &listQuantity)
	if err2 != nil {
		return err2
	}

	var ownby uint
	err3 := json.Unmarshal([]byte(data["ownby"]), &ownby)
	if err3 != nil {
		return err3
	}

	// {
	// 	"id": "[1,2,3]",
	// 	"quantity":[2,2,3]
	// 	"ownby":"123"
	// }
	for i := 0; i < len(listProductId); i++ {
		shopData := models.Shop{
			ProductId: listProductId[i],
			Quantity:  listQuantity[i],
			OwnBy:     ownby,
		}

		database.DB.Create(&shopData)
	}

	return c.JSON(fiber.Map{
		"message": "Successfully",
	})
}
