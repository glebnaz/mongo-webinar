package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type StoreController struct {
	cli *mongo.Client

	db   *mongo.Database
	coll *mongo.Collection
}

func (s *StoreController) FindById(ctx context.Context, id string, obj interface{}) error {
	filter := bson.M{"_id": id}
	err := s.coll.FindOne(ctx, filter).Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreController) Find(ctx context.Context, query bson.D, obj interface{}) error {
	cursor, err := s.coll.Find(ctx, query)
	if err != nil {
		return err
	}

	err = cursor.Decode(obj)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreController) UpdateMany(ctx context.Context, query bson.D, update bson.D) error {
	_, err := s.coll.UpdateMany(ctx, query, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreController) UpdateById(ctx context.Context, id string, update bson.D) error {
	filter := bson.M{"_id": id}
	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreController) DeleteMany(ctx context.Context, query bson.D) error {
	_, err := s.coll.DeleteMany(ctx, query)
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreController) DeleteById(ctx context.Context, id string) error {
	filter := bson.M{"_id": id}
	_, err := s.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
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
