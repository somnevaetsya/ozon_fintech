package usecases_impl

import (
	"ozon_test/app/models"
	"ozon_test/app/repositories"
	"ozon_test/app/usecases"
	"ozon_test/pkg/errors"
	"ozon_test/pkg/token"
)

type LinkUseCaseImpl struct {
	repLink repositories.LinkRepository
}

func MakeLinkUseCase(repLink_ repositories.LinkRepository) usecases.LinkUseCase {
	return &LinkUseCaseImpl{repLink: repLink_}
}

func (linkUseCase *LinkUseCaseImpl) GetLink(shortLink string) (models.Link, error) {
	link, err := linkUseCase.repLink.GetLink(shortLink)
	if err != nil {
		return models.Link{}, customErrors.ErrLinkNotFound
	}
	return *link, nil
}

func (linkUseCase *LinkUseCaseImpl) CreateLink(link *models.Link) error {
	if link.OriginalLink == "" {
		return customErrors.ErrBadInputData
	}
	isExist, err := linkUseCase.repLink.IsExistOriginal(link.OriginalLink)
	if err != nil && err != customErrors.ErrLinkNotFound {
		return err
	}
	if isExist {
		return customErrors.ErrLinkAlreadyExist
	}
	link.ShortLink = token.GenerateShortLink()
	isExistShort, err := linkUseCase.repLink.IsExistShort(link.ShortLink)
	if err != nil && err != customErrors.ErrLinkNotFound {
		return err
	}
	if isExistShort {
		for i := 0; i < 10; i++ {
			link.ShortLink = token.GenerateShortLink()
			isExistShort, err = linkUseCase.repLink.IsExistShort(link.ShortLink)
			if err != nil && err != customErrors.ErrLinkNotFound {
				return err
			}
			if !isExistShort {
				break
			}
		}
		if isExistShort {
			return customErrors.ErrLinkCouldNotBeCreated
		}
	}
	err = linkUseCase.repLink.CreateLink(link)
	if err != nil {
		return err
	}
	return nil
}
