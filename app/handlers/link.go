package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ozon_test/app/models"
	"ozon_test/app/usecases"
	"ozon_test/pkg/errors"
)

type LinkHandler struct {
	linkUseCase usecases.LinkUseCase
}

func MakeLinkHandler(linkUseCase_ usecases.LinkUseCase) *LinkHandler {
	return &LinkHandler{linkUseCase: linkUseCase_}
}

func (linkHandler *LinkHandler) CreateLink(c *gin.Context) {
	originalLink := c.Param("link")

	var link models.Link
	link.OriginalLink = originalLink
	err := linkHandler.linkUseCase.CreateLink(&link)
	if err != nil {
		_ = c.Error(err)
		return
	}

	linkJson, err := link.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusCreated, "application/json; charset=utf-8", linkJson)
}

func (linkHandler *LinkHandler) GetLink(c *gin.Context) {
	shortLink := c.Param("link")

	link, err := linkHandler.linkUseCase.GetLink(shortLink)
	if err != nil {
		_ = c.Error(err)
		return
	}

	linkJson, err := link.MarshalJSON()
	if err != nil {
		_ = c.Error(customErrors.ErrBadInputData)
		return
	}
	c.Data(http.StatusOK, "application/json; charset=utf-8", linkJson)
}
