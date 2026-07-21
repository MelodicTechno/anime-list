package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/MelodicTechno/anime-list/internal/model"
	"github.com/MelodicTechno/anime-list/internal/service"
)

type BookshelfHandler struct {
	svc *service.BookshelfService
}

func NewBookshelfHandler(svc *service.BookshelfService) *BookshelfHandler {
	return &BookshelfHandler{svc: svc}
}

type createBookshelfRequest struct {
	Name string `json:"name" binding:"required"`
}

type updateBookshelfRequest struct {
	Name string `json:"name" binding:"required"`
}

type addItemRequest struct {
	AnimeID int64  `json:"animeId" binding:"required"`
	StateID *int64 `json:"stateId"`
}

type bookshelfResponse struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"userId"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

func (h *BookshelfHandler) Create(c *gin.Context) {
	userID := c.GetInt64(ContextUserID)

	var req createBookshelfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bs, err := h.svc.Create(userID, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": toBookshelfResponse(bs)})
}

func (h *BookshelfHandler) List(c *gin.Context) {
	userID := c.GetInt64(ContextUserID)

	list, err := h.svc.List(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	results := make([]bookshelfResponse, len(list))
	for i, bs := range list {
		results[i] = toBookshelfResponse(&bs)
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}

func (h *BookshelfHandler) Get(c *gin.Context) {
	userID := c.GetInt64(ContextUserID)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	bs, items, err := h.svc.GetByID(userID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	type itemResponse struct {
		ID      int64  `json:"id"`
		AnimeID int64  `json:"animeId"`
		StateID *int64 `json:"stateId"`
	}

	itemList := make([]itemResponse, len(items))
	for i, it := range items {
		itemList[i] = itemResponse{ID: it.ID, AnimeID: it.AnimeID, StateID: it.StateID}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"bookshelf": toBookshelfResponse(bs),
			"items":     itemList,
		},
	})
}

func (h *BookshelfHandler) Update(c *gin.Context) {
	userID := c.GetInt64(ContextUserID)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req updateBookshelfRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bs, err := h.svc.Update(userID, id, req.Name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": toBookshelfResponse(bs)})
}

func (h *BookshelfHandler) Delete(c *gin.Context) {
	userID := c.GetInt64(ContextUserID)
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.svc.Delete(userID, id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nil})
}

func (h *BookshelfHandler) AddItem(c *gin.Context) {
	userID := c.GetInt64(ContextUserID)
	bookshelfID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req addItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.svc.AddItem(userID, bookshelfID, req.AnimeID, req.StateID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id":          item.ID,
			"bookshelfId": item.BookshelfID,
			"animeId":     item.AnimeID,
			"stateId":     item.StateID,
		},
	})
}

func (h *BookshelfHandler) RemoveItem(c *gin.Context) {
	userID := c.GetInt64(ContextUserID)
	bookshelfID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	itemID, err := strconv.ParseInt(c.Param("itemId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid item id"})
		return
	}

	if err := h.svc.RemoveItem(userID, bookshelfID, itemID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nil})
}

func toBookshelfResponse(bs *model.Bookshelf) bookshelfResponse {
	return bookshelfResponse{
		ID:        bs.ID,
		UserID:    bs.UserID,
		Name:      bs.Name,
		CreatedAt: bs.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

