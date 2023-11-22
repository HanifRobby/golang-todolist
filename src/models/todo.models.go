package models

import (
	"gorm.io/gorm"
)

// Defines todo tables for database communcications
type Todo struct {
	gorm.Model
	Name        string
	Description string
}
