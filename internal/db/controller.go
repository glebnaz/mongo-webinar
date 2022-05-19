package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	CurrencyDB         = "currency"
	CurrencyCollection = "currency"
)

type Controller struct {
	cli *mongo.Client

	db   *mongo.Database
	coll *mongo.Collection
}

func (c *Controller) InsertOne(ctx context.Context, obj interface{}) (id string, err error) {
	res, err := c.coll.InsertOne(ctx, obj)
	if err != nil {
		return
	}

	id, ok := res.InsertedID.(string)
	if !ok {
		return "", nil
	}
	return
}

func (c *Controller) InsertMany(ctx context.Context, obj []interface{}) (ids []string, err error) {
	res, err := c.coll.InsertMany(ctx, obj)
	if err != nil {
		return
	}

	for _, id := range res.InsertedIDs {
		ids = append(ids, id.(string))
	}
	return
}

func NewStoreController(cli *mongo.Client, db, coll string) Store {
	return &Controller{
		cli:  cli,
		db:   cli.Database(db),
		coll: cli.Database(db).Collection(coll),
	}
}
