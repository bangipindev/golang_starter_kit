package response

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

//
// ERROR STRUCT
//

type Error struct {
	Message  string `json:"message"`
	Code     string `json:"code"`
	HttpCode int    `json:"-"`
}

func NewError(msg string, code string, httpCode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

//
// DOMAIN ERRORS
//

var (

	// general
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbidden access")

	// auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 character")
	ErrAuthIsNotExists       = errors.New("auth is not exists")
	ErrEmailAlreadyUsed      = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")
	ErrRefreshTokenInvalid   = errors.New("refresh token is invalid")

	// product
	ErrProductRequired = errors.New("product is required")
	ErrProductInvalid  = errors.New("product must have minimum 4 character")
	ErrStockInvalid    = errors.New("stock must be greater than 0")
	ErrPriceInvalid    = errors.New("price must be greater than 0")

	// transaction
	ErrAmountInvalid          = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

//
// RESPONSE ERROR DEFINITIONS
//

var (
	ErrorGeneral         = NewError("Internal Server Error", "500", http.StatusInternalServerError)
	ErrorBadRequest      = NewError("bad request", "400", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), "404", http.StatusNotFound)
	ErrorUnauthorized    = NewError(ErrUnauthorized.Error(), "401", http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), "403", http.StatusForbidden)

	ErrorAuthIsNotExists     = NewError(ErrAuthIsNotExists.Error(), "404", http.StatusNotFound)
	ErrorEmailAlreadyUsed    = NewError(ErrEmailAlreadyUsed.Error(), "409", http.StatusConflict)
	ErrorPasswordNotMatch    = NewError(ErrPasswordNotMatch.Error(), "401", http.StatusUnauthorized)
	ErrorRefreshTokenInvalid = NewError(ErrRefreshTokenInvalid.Error(), "401", http.StatusUnauthorized)

	ErrorEmailRequired         = NewError(ErrEmailRequired.Error(), "400", http.StatusBadRequest)
	ErrorEmailInvalid          = NewError(ErrEmailInvalid.Error(), "400", http.StatusBadRequest)
	ErrorPasswordRequired      = NewError(ErrPasswordRequired.Error(), "400", http.StatusBadRequest)
	ErrorPasswordInvalidLength = NewError(ErrPasswordInvalidLength.Error(), "400", http.StatusBadRequest)
)

//
// ERROR MAPPING
//

var ErrorMapping = map[string]Error{
	ErrNotFound.Error():              ErrorNotFound,
	ErrUnauthorized.Error():          ErrorUnauthorized,
	ErrForbiddenAccess.Error():       ErrorForbiddenAccess,
	ErrEmailRequired.Error():         ErrorEmailRequired,
	ErrEmailInvalid.Error():          ErrorEmailInvalid,
	ErrPasswordRequired.Error():      ErrorPasswordRequired,
	ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
	ErrAuthIsNotExists.Error():       ErrorAuthIsNotExists,
	ErrEmailAlreadyUsed.Error():      ErrorEmailAlreadyUsed,
	ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,

	ErrRefreshTokenInvalid.Error(): ErrorRefreshTokenInvalid,
}

//
// HANDLE ERROR RESPONSE
//

func HandleError(c *fiber.Ctx, err error) error {

	if mappedErr, ok := ErrorMapping[err.Error()]; ok {

		return c.Status(mappedErr.HttpCode).JSON(BaseResponse{
			Success: false,
			Message: mappedErr.Message,
			Error: fiber.Map{
				"code": mappedErr.Code,
			},
		})
	}

	return c.Status(ErrorGeneral.HttpCode).JSON(BaseResponse{
		Success: false,
		Message: ErrorGeneral.Message,
		Error: fiber.Map{
			"code": ErrorGeneral.Code,
		},
	})
}
