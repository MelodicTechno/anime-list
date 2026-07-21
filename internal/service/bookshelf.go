package service

import (
	"errors"

	"github.com/MelodicTechno/anime-list/internal/model"
	"github.com/MelodicTechno/anime-list/internal/repository"
)

var (
	ErrBookshelfNotFound = errors.New("bookshelf not found")
	ErrBookshelfNotOwner = errors.New("bookshelf does not belong to current user")
	ErrItemNotFound      = errors.New("bookshelf item not found")
	ErrAnimeAlreadyAdded = errors.New("anime already exists in this bookshelf")
)

type BookshelfService struct {
	repo *repository.BookshelfRepository
}

func NewBookshelfService(repo *repository.BookshelfRepository) *BookshelfService {
	return &BookshelfService{repo: repo}
}

func (s *BookshelfService) Create(userID int64, name string) (*model.Bookshelf, error) {
	bs := &model.Bookshelf{
		UserID: userID,
		Name:   name,
	}
	if err := s.repo.Create(bs); err != nil {
		return nil, err
	}
	return bs, nil
}

func (s *BookshelfService) List(userID int64) ([]model.Bookshelf, error) {
	return s.repo.FindByUserID(userID)
}

func (s *BookshelfService) GetByID(userID, id int64) (*model.Bookshelf, []model.BookshelfItem, error) {
	bs, err := s.repo.FindByID(id)
	if err != nil {
		return nil, nil, err
	}
	if bs == nil {
		return nil, nil, ErrBookshelfNotFound
	}
	if bs.UserID != userID {
		return nil, nil, ErrBookshelfNotOwner
	}

	items, err := s.repo.FindItems(id)
	if err != nil {
		return nil, nil, err
	}

	return bs, items, nil
}

func (s *BookshelfService) Update(userID, id int64, name string) (*model.Bookshelf, error) {
	bs, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if bs == nil {
		return nil, ErrBookshelfNotFound
	}
	if bs.UserID != userID {
		return nil, ErrBookshelfNotOwner
	}

	bs.Name = name
	if err := s.repo.Update(bs); err != nil {
		return nil, err
	}
	return bs, nil
}

func (s *BookshelfService) Delete(userID, id int64) error {
	bs, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if bs == nil {
		return ErrBookshelfNotFound
	}
	if bs.UserID != userID {
		return ErrBookshelfNotOwner
	}

	return s.repo.Delete(id)
}

func (s *BookshelfService) AddItem(userID, bookshelfID, animeID int64, stateID *int64) (*model.BookshelfItem, error) {
	bs, err := s.repo.FindByID(bookshelfID)
	if err != nil {
		return nil, err
	}
	if bs == nil {
		return nil, ErrBookshelfNotFound
	}
	if bs.UserID != userID {
		return nil, ErrBookshelfNotOwner
	}

	exists, err := s.repo.ItemExists(bookshelfID, animeID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAnimeAlreadyAdded
	}

	item := &model.BookshelfItem{
		BookshelfID: bookshelfID,
		AnimeID:     animeID,
		StateID:     stateID,
	}
	if err := s.repo.AddItem(item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *BookshelfService) RemoveItem(userID, bookshelfID, itemID int64) error {
	bs, err := s.repo.FindByID(bookshelfID)
	if err != nil {
		return err
	}
	if bs == nil {
		return ErrBookshelfNotFound
	}
	if bs.UserID != userID {
		return ErrBookshelfNotOwner
	}

	item, err := s.repo.FindItemByID(itemID)
	if err != nil {
		return err
	}
	if item == nil || item.BookshelfID != bookshelfID {
		return ErrItemNotFound
	}

	return s.repo.RemoveItem(itemID)
}
