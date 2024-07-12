package entities

import "github.com/pborman/uuid"

/*
tweet => Struct privada
Tweet => Struct pública
*/
type Tweet struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func NewTweet() *Tweet { //*Tweet => referencia do endereço de memoria
	tweet := Tweet{
		ID: uuid.New(),
	}

	return &tweet //&tweet
}
