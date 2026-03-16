package response

import "github.com/gofiber/fiber/v2"

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessWithStatus(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
