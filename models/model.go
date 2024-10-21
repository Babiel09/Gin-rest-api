package models

type Amigo struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Nome      string `json:"nome"`
	Divertido bool   `json:"divertido"`
}
