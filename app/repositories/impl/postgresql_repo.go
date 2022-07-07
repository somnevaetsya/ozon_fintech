package repositories_impl

import (
	"github.com/jackc/pgx"
	"ozon_test/app/models"
	"ozon_test/app/repositories"
	"ozon_test/pkg/errors"
)

type PostgresRepository struct {
	db *pgx.ConnPool
}

func MakePostgresRepository(db *pgx.ConnPool) repositories.LinkRepository {
	return &PostgresRepository{db: db}
}

func (repo *PostgresRepository) CreateLink(link *models.Link) error {
	_, err := repo.db.Exec("insert into links values ($1, $2);", link.OriginalLink, link.ShortLink)
	return err
}
func (repo *PostgresRepository) GetLink(shortLink string) (*models.Link, error) {
	link := new(models.Link)
	err := repo.db.QueryRow("select original_link, short_link from links where short_link = $1;", shortLink).Scan(&link.OriginalLink, &link.ShortLink)
	return link, err
}

func (repo *PostgresRepository) IsExistOriginal(originalLink string) (bool, error) {
	link := new(models.Link)
	err := repo.db.QueryRow("select original_link, short_link from links where original_link = $1;", originalLink).Scan(&link.OriginalLink, &link.ShortLink)
	if err != nil {
		return false, customErrors.ErrLinkNotFound
	}
	if link.OriginalLink != "" {
		return true, nil
	} else {
		return false, nil
	}
}

func (repo *PostgresRepository) IsExistShort(shortLink string) (bool, error) {
	link := new(models.Link)
	err := repo.db.QueryRow("select original_link, short_link from links where short_link = $1;", shortLink).Scan(&link.OriginalLink, &link.ShortLink)
	if err != nil {
		return false, customErrors.ErrLinkNotFound
	}
	if link.OriginalLink != "" {
		return true, nil
	} else {
		return false, nil
	}
}
