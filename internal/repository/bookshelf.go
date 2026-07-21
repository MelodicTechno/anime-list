package repository

import (
	"github.com/MelodicTechno/anime-list/internal/model"
	"gorm.io/gorm"
)

type BookshelfRepository struct {
	db *gorm.DB
}

func NewBookshelfRepository(db *gorm.DB) *BookshelfRepository {
	return &BookshelfRepository{db: db}
}

func (r *BookshelfRepository) Create(bs *model.Bookshelf) error {
	return r.db.Create(bs).Error
}

func (r *BookshelfRepository) FindByID(id int64) (*model.Bookshelf, error) {
	var bs model.Bookshelf
	err := r.db.First(&bs, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &bs, err
}

func (r *BookshelfRepository) FindByUserID(userID int64) ([]model.Bookshelf, error) {
	var list []model.Bookshelf
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&list).Error
	return list, err
}

func (r *BookshelfRepository) Update(bs *model.Bookshelf) error {
	return r.db.Save(bs).Error
}

func (r *BookshelfRepository) Delete(id int64) error {
	return r.db.Delete(&model.Bookshelf{}, id).Error
}

func (r *BookshelfRepository) AddItem(item *model.BookshelfItem) error {
	return r.db.Create(item).Error
}

func (r *BookshelfRepository) RemoveItem(itemID int64) error {
	return r.db.Delete(&model.BookshelfItem{}, itemID).Error
}

func (r *BookshelfRepository) FindItemByID(itemID int64) (*model.BookshelfItem, error) {
	var item model.BookshelfItem
	err := r.db.First(&item, itemID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &item, err
}

func (r *BookshelfRepository) FindItems(bookshelfID int64) ([]model.BookshelfItem, error) {
	var items []model.BookshelfItem
	err := r.db.Where("bookshelf_id = ?", bookshelfID).Find(&items).Error
	return items, err
}

func (r *BookshelfRepository) ItemExists(bookshelfID, animeID int64) (bool, error) {
	var count int64
	err := r.db.Model(&model.BookshelfItem{}).
		Where("bookshelf_id = ? AND anime_id = ?", bookshelfID, animeID).
		Count(&count).Error
	return count > 0, err
}
