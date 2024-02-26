package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Password    string             `json:"-"`
	APIKey      string             `json:"apiKey,omitempty" bson:"apiKey,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Twitter     string             `json:"twitter,omitempty" bson:"twitter,omitempty"`
	DiscordData string             `json:"discordData,omitempty" bson:"discordData,omitempty"`
	Email       string             `json:"email,omitempty" bson:"email"`
	DateCreated time.Time          `json:"dateCreated" bson:"dateCreated"`
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
}
