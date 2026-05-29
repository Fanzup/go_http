package link

import (
	"demo/weather_check/packages/db"

	"gorm.io/gorm/clause"
)

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

func (repo *LinkRepository) Update(link *Link) (*Link, error) {
	res := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (repo *LinkRepository) Delete(id uint) error {
	res := repo.DataBase.DB.Delete(&Link{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (repo *LinkRepository) GetById(id uint) (*Link, error) {
	var link Link
	res := repo.DataBase.DB.First(&link, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}
