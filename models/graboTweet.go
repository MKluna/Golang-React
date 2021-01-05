package models

import "time"

/*GraboTweet es el formato o estructura que tendra nuestro tweet*/
type GraboTweet struct {
	UserID string    `bson:"userid" json:"userid,omitempty"`
	Menaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha  time.Time `bson:"fecha" fecha:"userid,omitempty"`
}
