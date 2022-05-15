package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
)

type Store interface {
	InsertOne(ctx context.Context, obj interface{}) (id string, err error)
	InsertMany(ctx context.Context, obj []interface{}) (ids []string, err error)
	FindById(ctx context.Context, id string, obj interface{}) error
	Find(ctx context.Context, query bson.D, obj interface{}) error
	UpdateMany(ctx context.Context, query bson.D, update bson.D) error
	UpdateById(ctx context.Context, id string, update bson.D) error
	DeleteMany(ctx context.Context, query bson.D) error
	DeleteById(ctx context.Context, id string) error
}
