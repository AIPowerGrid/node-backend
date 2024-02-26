package db

import (
	"backend/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUser(u models.User) (primitive.ObjectID, error) {
	col := _getCol("users")
	res, err := col.InsertOne(_ctx(), u)
	_s := res.InsertedID.(primitive.ObjectID)
	return _s, err
	// id := _s.Hex()
	// log.Info(id)
	// return id,err
}
