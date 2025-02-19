package http

// import (
// 	// "strconv"

// 	"github.com/signshine/give-a-sign/api/service"

// 	"github.com/gofiber/fiber/v2"
// )

// type WordHandler struct {
// 	svc *service.WordService
// }

// func NewWordHandler(service *service.WordService) *WordHandler {
// 	return &WordHandler{
// 		svc: service,
// 	}
// }

// func (h *WordHandler) AddWord(c *fiber.Ctx) error {
// 	var req presenter.Word

// 	if err := c.BodyParser(&req); err != nil {
// 		return fiber.ErrBadRequest
// 	}

// 	word := presenter.WordPresenter2Domain(&req)
// 	resp, err := h.svc.CreateWord(c.UserContext(), *word)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(map[string]interface{}{
// 		"message": "success",
// 		"id":      resp,
// 	})
// }

// func (h *WordHandler) AllWords(c *fiber.Ctx) error {
// 	page, err := c.ParamsInt("page")
// 	if err != nil {
// 		return fiber.ErrBadRequest
// 	}
// 	_ = page

// 	pageSize, err := c.ParamsInt("pageSize")
// 	if err != nil {
// 		return fiber.ErrBadRequest
// 	}
// 	_ = pageSize

// 	panic("unimplemented")
// }

// func (h *WordHandler) GetWordById(c *fiber.Ctx) error {
// 	panic("unimplemented")
// }

// func (h *WordHandler) AddVideo(c *fiber.Ctx) error {
// 	var req presenter.Video

// 	if err := c.BodyParser(&req); err != nil {
// 		return fiber.ErrBadRequest
// 	}

// 	wordId, err := strconv.Atoi(c.Params("wordId"))
// 	if wordId == 0 && err != nil {
// 		return fiber.ErrBadRequest
// 	}
// 	req.WordID = uint(wordId)

// 	video := presenter.VideoPresenter2Domain(&req)
// 	resp, err := h.svc.CreateVideo(c.UserContext(), *video)
// 	if err != nil {
// 		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
// 	}

// 	return c.JSON(map[string]interface{}{
// 		"message": "success",
// 		"videoId": resp,
// 	})
// }

// func (h *WordHandler) AllVideos(c *fiber.Ctx) error {
// 	panic("unimplemented")
// }

// func (h *WordHandler) GetVideoById(c *fiber.Ctx) error {
// 	panic("unimplemented")
// }
