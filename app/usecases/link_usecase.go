package usecases

import "ozon_test/app/models"

type LinkUseCase interface {
	GetLink(link string) (models.Link, error)
	CreateLink(link *models.Link) error
}
