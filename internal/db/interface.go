package db

import "context"

type Store interface {
	InsertOne(ctx context.Context, obj interface{}) (id string, err error)
	InsertMany(ctx context.Context, obj []interface{}) (ids []string, err error)
}
