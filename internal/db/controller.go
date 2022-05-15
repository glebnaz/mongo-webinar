package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type StoreController struct {
	cli *mongo.Client

	db   *mongo.Database
	coll *mongo.Collection
}

func NewStoreController(cli *mongo.Client, db string, collections string) Store {
	return &StoreController{
		cli:  cli,
		db:   cli.Database(db),
		coll: cli.Database(db).Collection(collections),
	}
}

func (s *StoreController) InsertOne(ctx context.Context, obj interface{}) (id string, err error) {
	res, err := s.coll.InsertOne(ctx, obj)
	if err != nil {
		log.Printf("Error inserting %v", err)
		return "", err
	}

	id = fmt.Sprintf("%v", res.InsertedID)
	return
}

func (s *StoreController) InsertMany(ctx context.Context, obj []interface{}) (ids []string, err error) {
	res, err := s.coll.InsertMany(ctx, obj)
	if err != nil {
		log.Printf("Error inserting many%v", err)
		return nil, err
	}

	ids = make([]string, len(res.InsertedIDs), 0)

	for _, id := range res.InsertedIDs {
		ids = append(ids, fmt.Sprintf("%v", id))
	}
	return
}
