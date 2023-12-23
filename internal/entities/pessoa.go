package entities

import "time"

type Pessoa struct {
	ID 		string 		`json:"id"`
	Nome		string 		`json:"nome"`
	Apelido		string		`json:"apelido"`
	Nascimento 	time.Time	`json:"nascimento"`
	Stack	 	string		`json:"string"`
}
