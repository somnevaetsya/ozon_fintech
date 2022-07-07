package repositories

import "ozon_test/app/models"

type LinkRepository interface {
	CreateLink(link *models.Link) error
	GetLink(shortLink string) (*models.Link, error)
	IsExistOriginal(originalLink string) (bool, error)
	IsExistShort(shortLink string) (bool, error)
}
