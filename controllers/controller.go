package controllers

import (
	"yofio/avanzado/store"

	"gorm.io/gorm"
)

// A Controller response
type Controller struct {
	repo store.IRepository
}

// NewController function
func NewController(db *gorm.DB) *Controller {
	return &Controller{
		repo: store.NewRepository(db),
	}
}
