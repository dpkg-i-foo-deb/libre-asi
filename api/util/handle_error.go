package util

import (
	"libre-asi-api/errors"
	"libre-asi-api/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func HandleErorLog(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HandleErrorStop(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func HandleFiberError(c *fiber.Ctx, err error) error {

	status := 200
	var res models.Response

	switch err {
	case errors.ErrInvalidCredentials:
		status = 401
		res.Status = string(models.DENIED)
		res.Message = "Invalid Credentials"
	case errors.ErrSetupRequired:
		status = 412
		res.Status = string(models.SETUP_REQUIRED)
		res.Message = "Libre-ASI Requires Set-Up"
	case errors.ErrCheckRequest:
		status = 400
		res.Status = string(models.CHECK_REQUEST)
		res.Message = "Check your request"
	case errors.ErrNotImplemmented:
		status = 501
		res.Status = string(models.ERROR)
		res.Message = "Functionality is not yet implemmented"
	case errors.ErrNoData:
		status = 400
		res.Status = string(models.CHECK_REQUEST)
		res.Message = "Not enough data was present"
	case errors.ErrInternalError:
		status = 500
		res.Status = string(models.ERROR)
		res.Message = "Something went wrong on the server side"
	case errors.ErrTooManyEntities:
		status = 413
		res.Status = string(models.DENIED)
		res.Message = "Too many entities were present"
	case errors.ErrBadRoute:
		status = 400
		res.Status = string(models.DENIED)
		res.Message = "Bad route"
	case errors.ErrAccessDenied:
		status = 401
		res.Status = string(models.DENIED)
		res.Message = "Access denied, not enough privileges"
	case errors.ErrConflict:
		status = 409
		res.Status = string(models.DENIED)
		res.Message = "The entity already exists"
	}

	return c.Status(status).JSON(res)
}
