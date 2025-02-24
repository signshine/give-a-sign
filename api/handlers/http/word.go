package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/signshine/give-a-sign/api/pb"
	"github.com/signshine/give-a-sign/api/service"
	"github.com/signshine/give-a-sign/pkg/context"
)

func CreateWord(svcGetter serviceGetter[*service.WordService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.CreateWordRequest

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		resp, err := svc.CreateWord(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrWordAlreadyExist) {
				return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
			}
			if errors.Is(err, service.ErrWordOnCreate) {
				return fiber.ErrBadRequest
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetWord(svcGetter serviceGetter[*service.WordService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		var req pb.GetWordRequest
		var filter pb.WordFilter

		errBody := c.BodyParser(&req)
		errQuery := c.QueryParser(&filter)
		if errBody != nil && errQuery != nil {
			return fiber.ErrBadRequest
		}

		if req.Filter == nil {
			req.Filter = &filter
		}

		resp, err := svc.GetWord(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrWordNotFound) {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			if errors.Is(err, service.ErrWordFilterValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(resp)
	}
}

func GetListWord(svcGetter serviceGetter[*service.WordService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		var req pb.ListWordRequest

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.GetAllWords(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrPaginationNegativePage) || 
			errors.Is(err, service.ErrPaginationNegativePagesize) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(resp)
	}
}

func DeleteWord(svcGetter serviceGetter[*service.WordService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		var req pb.DeleteWordRequest
		var filter pb.WordFilter

		errBody := c.BodyParser(&req)
		errQuery := c.QueryParser(&filter)
		if errBody != nil && errQuery != nil {
			return fiber.ErrBadRequest
		}

		if req.Filter == nil {
			req.Filter = &filter
		}

		resp, err := svc.DeleteWord(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrWordNotFound) {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			if errors.Is(err, service.ErrWordFilterValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(resp)
	}
}