package repositories_impl

import (
	"ozon_test/app/models"
	"ozon_test/app/repositories"
	customErrors "ozon_test/pkg/errors"
)

type InMemoryRepository struct {
	OriginalValues map[string]string
	ShortValues    map[string]string
}

func MakeInMemoryRepository() repositories.LinkRepository {
	return &InMemoryRepository{OriginalValues: make(map[string]string), ShortValues: make(map[string]string)}
}

func (repo *InMemoryRepository) CreateLink(link *models.Link) error {
	repo.ShortValues[link.ShortLink] = link.OriginalLink
	repo.OriginalValues[link.OriginalLink] = link.ShortLink
	return nil
}
func (repo *InMemoryRepository) GetLink(shortLink string) (*models.Link, error) {
	if link, ok := repo.ShortValues[shortLink]; ok {
		return &models.Link{OriginalLink: link, ShortLink: shortLink}, nil
	} else {
		return nil, customErrors.ErrLinkNotFound
	}
}
func (repo *InMemoryRepository) IsExistOriginal(originalLink string) (bool, error) {
	if _, ok := repo.OriginalValues[originalLink]; ok {
		return true, nil
	} else {
		return false, nil
	}
}

func (repo *InMemoryRepository) IsExistShort(shortLink string) (bool, error) {
	if _, ok := repo.ShortValues[shortLink]; ok {
		return true, nil
	} else {
		return false, nil
	}
}
