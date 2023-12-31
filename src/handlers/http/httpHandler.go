package http

import (
	"github.com/SyaibanAhmadRamadhan/api-pagination-golang/src/modules/profile/services"
)

type HttpHandlerImpl struct {
	ProfileService services.ProfileService
}

func NewHttpHandlerImpl(
	profileService services.ProfileService,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		ProfileService: profileService,
	}
}

func (hand *HttpHandlerImpl) HttpHandlerRouter(r *fiber.App) {
	r.Route("/api/v1", func(router fiber.Router) {
		router.Get("/profile", hand.FindAll)
		router.Post("/profile", hand.InsertProfile)
		router.Get("/profile/:id", hand.FindByID)
		router.Put("/profile/:id", hand.UpdateProfile)
		router.Delete("/profile/:id", hand.DeleteById)
	})
}
