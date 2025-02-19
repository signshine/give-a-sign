package http

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/signshine/give-a-sign/api/pb"
	"github.com/signshine/give-a-sign/api/service"
	"github.com/signshine/give-a-sign/pkg/context"
)

func CreateLanguage(svcGetter serviceGetter[*service.LanguageService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.CreateLanguageRequest

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		resp, err := svc.CreateLanguage(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrLanguageAlreadyExist) {
				return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
			}
			if errors.Is(err, service.ErrLanguageOnCreate) {
				return fiber.ErrBadRequest
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func GetLanguage(svcGetter serviceGetter[*service.LanguageService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		var req pb.GetLanguageRequest
		var filter pb.LanguageFilter

		errBody := c.BodyParser(&req)
		errQuery := c.QueryParser(&filter)
		if errBody != nil && errQuery != nil {
			return fiber.ErrBadRequest
		}

		if req.Filter == nil {
			req.Filter = &filter
		}

		resp, err := svc.GetLanguage(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrLanguageNotFound) {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			if errors.Is(err, service.ErrLanguageFilterValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(resp)
	}
}

func GetListLanguage(svcGetter serviceGetter[*service.LanguageService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		var req pb.ListLanguagesRequest

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		resp, err := svc.GetAllLanguage(c.UserContext(), &req)
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

func DeleteLanguage(svcGetter serviceGetter[*service.LanguageService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		logger := context.GetLogger(c.UserContext())
		svc := svcGetter(c.UserContext())

		var req pb.DeleteLanguageRequest
		var filter pb.LanguageFilter

		errBody := c.BodyParser(&req)
		errQuery := c.QueryParser(&filter)
		if errBody != nil && errQuery != nil {
			return fiber.ErrBadRequest
		}

		if req.Filter == nil {
			req.Filter = &filter
		}

		resp, err := svc.DeleteLanguage(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrLanguageNotFound) {
				return fiber.NewError(fiber.StatusNotFound, err.Error())
			}
			if errors.Is(err, service.ErrLanguageFilterValidation) {
				return fiber.NewError(fiber.StatusBadRequest, err.Error())
			}

			return fiber.ErrInternalServerError
		}

		return c.JSON(resp)
	}
}