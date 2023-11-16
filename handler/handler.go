package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nyaruka/phonenumbers"
	
	"github.com/radianhanggata/siesta-coding-test/q3/model"
)

func FormatPhoneNumber(c *fiber.Ctx) error {
	phonenumber := c.Params("phonenumber")
	countryCode := c.Params("countrycode")
	num, err := phonenumbers.Parse(phonenumber, countryCode)
	if err != nil {
		return c.Status(400).Send([]byte("Bad Request"))
	}

	if !phonenumbers.IsValidNumber(num) {
		return c.Status(400).Send([]byte("Bad Request"))
	}

	formatted := phonenumbers.Format(num, phonenumbers.E164)
	response := model.FormatPhoneNumberResponse{PhoneNumber: formatted[1:]}

	return c.Status(200).JSON(response)
}
