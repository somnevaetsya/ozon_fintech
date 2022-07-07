package usecases_impl

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"ozon_test/app/models"
	"ozon_test/app/repositories/mock"
	"ozon_test/pkg/errors"
	"testing"
)

func TestGetLink(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	linkRepo := mock_repositories.NewMockLinkRepository(controller)
	linkUseCase := MakeLinkUseCase(linkRepo)
	link := models.Link{ShortLink: "aa"}
	linkRepo.EXPECT().GetLink("aa").Return(&link, nil)
	expectedLink, err := linkUseCase.GetLink("aa")
	assert.Equal(t, link, expectedLink)
	assert.Equal(t, err, nil)

	linkRepo.EXPECT().GetLink("aao").Return(&models.Link{}, customErrors.ErrLinkNotFound)
	expectedLink, err = linkUseCase.GetLink("aao")
	assert.Equal(t, models.Link{}, expectedLink)
	assert.Equal(t, err, customErrors.ErrLinkNotFound)
}

func TestCreateLink(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	linkRepo := mock_repositories.NewMockLinkRepository(controller)
	linkUseCase := MakeLinkUseCase(linkRepo)
	link := models.Link{OriginalLink: ""}
	err := linkUseCase.CreateLink(&link)
	assert.Equal(t, err, customErrors.ErrBadInputData)

	link2 := models.Link{OriginalLink: "aao"}
	linkRepo.EXPECT().IsExistOriginal("aao").Return(false, customErrors.ErrLinkAlreadyExist)
	err = linkUseCase.CreateLink(&link2)
	assert.Equal(t, err, customErrors.ErrLinkAlreadyExist)

	linkRepo.EXPECT().IsExistOriginal("aao").Return(true, nil)
	err = linkUseCase.CreateLink(&link2)
	assert.Equal(t, err, customErrors.ErrLinkAlreadyExist)

	linkRepo.EXPECT().IsExistOriginal("aao").Return(false, nil)
	linkRepo.EXPECT().IsExistShort(gomock.Any()).Return(false, customErrors.ErrLinkAlreadyExist)
	err = linkUseCase.CreateLink(&link2)
	assert.Equal(t, err, customErrors.ErrLinkAlreadyExist)

	linkRepo.EXPECT().IsExistOriginal("aao").Return(false, nil)
	linkRepo.EXPECT().IsExistShort(gomock.Any()).Return(false, nil)
	linkRepo.EXPECT().CreateLink(&link2)
	err = linkUseCase.CreateLink(&link2)
	assert.Equal(t, err, nil)
}
