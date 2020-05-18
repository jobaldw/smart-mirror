package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Elements additional fields
type Elements struct {
	ID *primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`

	Name      string    `bson:"name,omitempty" json:"name,omitempty"`
	Genres    []string  `bson:"genres,omitempty" json:"genres,omitempty"`
	Platforms []string  `bson:"platforms,omitempty" json:"platforms,omitempty"`
	Runtime   string    `bson:"runtime,omitempty" json:"runtime,omitempty"`
	Created   time.Time `bson:"created,omitempty" json:"created,omitempty"`
	Updated   time.Time `bson:"updated,omitempty" json:"updated,omitempty"`
	Watched   bool      `bson:"watched,omitempty" json:"watched,omitempty"`
}

//Movie model
type Movie struct {
	Elements `bson:",inline" json:",inline"`

	Franchise string `bson:"franchise,omitempty" json:"franchise,omitempty"`
}

//Show model
type Show struct {
	Elements `bson:",inline" json:",inline"`

	Review  string `bson:"review,omitempty" json:"review,omitempty"`
	Airing  bool   `bson:"airing,omitempty" json:"airing,omitempty"`
	Seasons int    `bson:"seasons,omitempty" json:"seasons,omitempty"`
}