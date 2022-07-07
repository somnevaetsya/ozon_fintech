package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"ozon_test/app/middleware"
	"ozon_test/app/models"
	mock_usecases "ozon_test/app/usecases/mock"
	customErrors "ozon_test/pkg/errors"
	"testing"
)

func TestGetAttachment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	linkUseCase := mock_usecases.NewMockLinkUseCase(controller)
	linkHandler := MakeLinkHandler(linkUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	mainRoutes := router.Group("/")
	{
		mainRoutes.GET("/:link", linkHandler.GetLink)
	}

	link := models.Link{ShortLink: "link"}

	//good
	linkUseCase.EXPECT().GetLink("link").Return(link, nil)
	request, _ := http.NewRequest("GET", "/link", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusOK, writer.Code)
	var expectedLink models.Link
	_ = json.Unmarshal(writer.Body.Bytes(), &expectedLink)
	assert.Equal(t, link, expectedLink)

	//bad
	linkUseCase.EXPECT().GetLink("link").Return(link, customErrors.ErrLinkNotFound)
	request, _ = http.NewRequest("GET", "/link", nil)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusNotFound, writer.Code)
}

func TestCreateAttachment(t *testing.T) {
	t.Parallel()
	controller := gomock.NewController(t)
	defer controller.Finish()
	linkUseCase := mock_usecases.NewMockLinkUseCase(controller)
	linkHandler := MakeLinkHandler(linkUseCase)

	router := gin.Default()
	router.Use(middleware.CheckError())

	mainRoutes := router.Group("/")
	{
		mainRoutes.POST("/:link", linkHandler.CreateLink)
	}

	emptyLink := models.Link{OriginalLink: "aa"}

	//good
	linkUseCase.EXPECT().CreateLink(&emptyLink).Return(customErrors.ErrBadInputData)
	request, _ := http.NewRequest("POST", "/aa", nil)
	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusBadRequest, writer.Code)

	//bad
	linkUseCase.EXPECT().CreateLink(&emptyLink).Return(nil)
	request, _ = http.NewRequest("POST", "/aa", nil)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, request)
	assert.Equal(t, http.StatusCreated, writer.Code)
	var expectedLink models.Link
	_ = json.Unmarshal(writer.Body.Bytes(), &expectedLink)
	assert.Equal(t, emptyLink, expectedLink)
}
