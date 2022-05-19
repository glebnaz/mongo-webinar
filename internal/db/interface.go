package db

import (
	"context"
	"github.com/glebnaz/mongo-webinar/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type Store interface {
	InsertOne(ctx context.Context, obj interface{}) (id string, err error)
	InsertMany(ctx context.Context, obj []interface{}) (ids []string, err error)
	Find(ctx context.Context, filter bson.D) (curr []model.Currency, err error)
}
