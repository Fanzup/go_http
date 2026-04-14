package link

import "demo/weather_check/packages/db"

type LinkRepository struct {
	DataBase *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		DataBase: database,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	res := repo.DataBase.DB.Create(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	res := repo.DataBase.DB.First(&link, "hash = ?", hash)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}
