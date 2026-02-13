package response

import "github.com/gofiber/fiber/v2"

type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, status int, message string, err interface{}) error {
	return c.Status(status).JSON(BaseResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}
