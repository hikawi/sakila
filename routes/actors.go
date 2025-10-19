package routes

import (
	"context"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
	"luny.dev/awad-w1/models"
)

type ActorsHandler struct {
	DB *gorm.DB
}

type GetActorsQuery struct {
	Page    int `form:"page,default=1" binding:"omitempty,min=1"`
	PerPage int `form:"per_page,default=20" binding:"omitempty,min=1,max=100"`
}

type PostActorBody struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type PatchActorBody struct {
	FirstName string `json:"first_name" binding:"omitempty"`
	LastName  string `json:"last_name" binding:"omitempty"`
}

func (h *ActorsHandler) GetActors(c *gin.Context) {
	_, err := h.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't connect to database"})
		return
	}

	var query GetActorsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	res, err := gorm.G[models.Actor](h.DB).
		Limit(query.PerPage).Offset((query.Page - 1) * query.PerPage).Find(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	count, err := gorm.G[models.Actor](h.DB).Count(ctx, "actor_id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        res,
		"page":        query.Page,
		"per_page":    query.PerPage,
		"total":       count,
		"total_pages": int64(math.Ceil(float64(count) / float64(query.PerPage))),
	})
}

func (h *ActorsHandler) GetActorID(c *gin.Context) {
	_, err := h.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't connect to database"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a non-negative number"})
		return
	}

	ctx := context.Background()
	res, err := gorm.G[models.Actor](h.DB).Where("actor_id = ?", id).First(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "that actor can not be found"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *ActorsHandler) PostActor(c *gin.Context) {
	_, err := h.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't connect to database"})
		return
	}

	var body PostActorBody
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	actor := models.Actor{
		FirstName: body.FirstName,
		LastName:  body.LastName,
	}

	ctx := context.Background()
	err = gorm.G[models.Actor](h.DB).Create(ctx, &actor)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func (h *ActorsHandler) PatchActor(c *gin.Context) {
	_, err := h.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't connect to database"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a non-negative number"})
		return
	}

	var body PatchActorBody
	if err := c.ShouldBindWith(&body, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := context.Background()
	res, err := gorm.G[models.Actor](h.DB).Where("actor_id = ?", id).Updates(ctx, models.Actor{FirstName: body.FirstName, LastName: body.LastName})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"affected": res})
}

func (h *ActorsHandler) DeleteActor(c *gin.Context) {
	_, err := h.DB.DB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't connect to database"})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 16)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a non-negative number"})
		return
	}

	ctx := context.Background()
	affected, err := gorm.G[models.Actor](h.DB).Where("actor_id = ?", id).Delete(ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"affected": affected})
}
