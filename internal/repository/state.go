package repository

import (
	"github.com/MelodicTechno/anime-list/internal/model"
	"gorm.io/gorm"
)

type StateRepository struct {
	db *gorm.DB
}

func NewStateRepository(db *gorm.DB) *StateRepository {
	return &StateRepository{db: db}
}

func (r *StateRepository) FindAll() ([]model.State, error) {
	var states []model.State
	err := r.db.Order("id").Find(&states).Error
	return states, err
}

func (r *StateRepository) FindByID(id int64) (*model.State, error) {
	var state model.State
	err := r.db.First(&state, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &state, err
}
