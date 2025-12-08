package model

type Jogador struct {
	ID     int    `json:"id"`
	Nome   string `json:"nome"`
	Idade  int    `json:"idade"`
	TimeID int    `json:"time_id"`
}
