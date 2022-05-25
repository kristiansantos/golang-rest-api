package models

import "gorm.io/gorm"

type Planet struct {
	gorm.Model
	Name            string `json:"name" binding:"required"`
	Climate         string `json:"climare" binding:"required"`
	Ground          string `json:"ground" binding:"required"`
	FilmAppearances int    `json:"film_appearances"`
}
