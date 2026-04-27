package user

import "github.com/gofiber/fiber/v3"

type ProfileResponse struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (r *Router) Profile() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		userId := ctx.Locals("userId").(int)

		profile, err := r.service.Profile(ctx, userId)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}

		return ctx.Status(fiber.StatusOK).JSON(ProfileResponse{
			Id:        profile.Id,
			FirstName: profile.FirstName,
			LastName:  profile.LastName,
		})
	}
}
