package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name,omitempty"`
	Description   string             `bson:"description,omitempty"`
	Director      string             `bson:"director,omitempty"`
	Actors        []string           `bson:"actors,omitempty"`
	HeroName      string             `bson:"hero_name,omitempty"`
	HeroinName    string             `bson:"heroin_name,omitempty"`
	YearOfRelease string             `bson:"year_of_release,omitempty"`
}
