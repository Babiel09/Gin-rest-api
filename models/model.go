package models

type Amigo struct {
	Id        int    `json:"id"`
	Nome      string `json:"nome"`
	Divertido bool   `json:"divertido"`
}

//Definindo o slice:

var Amigos []Amigo
