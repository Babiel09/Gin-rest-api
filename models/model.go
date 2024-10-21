package models

import "gorm.io/gorm"

type Amigo struct {
	gorm.Model
	Nome      string `json:"nome"`
	Divertido bool   `json:"divertido"`
}
