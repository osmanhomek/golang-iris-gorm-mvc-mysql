package repos

import (
	"gomvc/models"

	"github.com/jinzhu/gorm"
)

type BookRepository interface {
	Select(query string) []models.Book
	SelectById(query string, id int64) models.Book
	SelectByName(query string, name string) models.Book
}

type bookMysqlRepository struct {
	DB *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookMysqlRepository{DB: db}
}

func (m *bookMysqlRepository) Select(query string) []models.Book {
	result := []models.Book{}
	m.DB.Raw(query).Scan(&result)
	return result
}

func (m *bookMysqlRepository) SelectById(query string, id int64) models.Book {
	result := models.Book{}
	m.DB.Raw(query, id).Scan(&result)
	return result
}

func (m *bookMysqlRepository) SelectByName(query string, name string) models.Book {
	result := models.Book{}
	m.DB.Raw(query, name).Scan(&result)
	return result
}
