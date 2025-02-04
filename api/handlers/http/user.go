package http

import (
	"errors"

	"github.com/signshine/give-a-sign/api/pb"
	"github.com/signshine/give-a-sign/api/service"
	"github.com/signshine/give-a-sign/pkg/context"

	"github.com/gofiber/fiber/v2"
)

func SignUp(svcGetter serviceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.UserSignUpRequest

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		svc := svcGetter(c.UserContext())
		logger := context.GetLogger(c.UserContext())

		resp, err := svc.SignUp(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrUserCreationValidation) {
				return fiber.ErrBadRequest
			}
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}

func SignIn(svcGetter serviceGetter[*service.UserService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req pb.UserSignInRequest

		if err := c.BodyParser(&req); err != nil {
			return fiber.ErrBadRequest
		}

		svc := svcGetter(c.UserContext())
		logger := context.GetLogger(c.UserContext())

		resp, err := svc.SignIn(c.UserContext(), &req)
		if err != nil {
			logger.Error(err.Error())

			if errors.Is(err, service.ErrUserFilterValidation) {
				return fiber.ErrBadRequest
			}
			if errors.Is(err, service.ErrUserWrongPassword) ||
				errors.Is(err, service.ErrUserNotFound) {
				return fiber.NewError(fiber.StatusUnauthorized, err.Error())
			}

			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return c.JSON(resp)
	}
}
