package response

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func ValidationError(c *fiber.Ctx, err error) error {

	var errors = make(map[string]string)

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		for _, e := range validationErrors {

			switch e.Tag() {

			case "required":
				errors[e.Field()] = e.Field() + " is required"

			case "email":
				errors[e.Field()] = "invalid email format"

			case "min":
				errors[e.Field()] = e.Field() + " minimum length is " + e.Param()

			case "max":
				errors[e.Field()] = e.Field() + " maximum length is " + e.Param()

			default:
				errors[e.Field()] = "invalid value"
			}
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(BaseResponse{
		Success: false,
		Message: "validation failed",
		Error:   errors,
	})
}
