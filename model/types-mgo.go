// +build mgo

package model

import (
	"errors"
	"time"

	"github.com/globalsign/mgo/bson"
)

type ObjectId = bson.ObjectId
type AccountId = ObjectId
type Raw = bson.Raw

var(
	ErrNoDocuments = errors.New("mongo: no documents in result")
)

func NewObjectId()ObjectId {
	return bson.NewObjectId()
}

func GetTime(id ObjectId)time.Time {
	return id.Time()
}