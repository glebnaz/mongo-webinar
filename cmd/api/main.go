package main

import (
	"github.com/glebnaz/mongo-webinar/internal/db"
	"github.com/glebnaz/mongo-webinar/internal/model"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ReqGet struct {
	DateStart int64  `json:"date_start"`
	DateEnd   int64  `json:"date_end"`
	CharCode  string `json:"char_code"`
}

func main() {
	e := echo.New()

	cli, err := mongo.NewClient()
	if err != nil {
		panic(err)
	}

	err = cli.Connect(nil)
	if err != nil {
		panic(err)
	}

	store := db.NewStoreController(cli, db.CurrencyDB, db.CurrencyCollection)

	e.POST("/get", func(c echo.Context) error {
		var req ReqGet

		if err := c.Bind(&req); err != nil {
			return err
		}

		filter := bson.D{
			{"date", bson.D{
				{"$gte", req.DateStart},
				{"$lte", req.DateEnd},
			}},
			{"char_code", req.CharCode},
		}

		var curencies []model.Currency

		if err := store.Find(c.Request().Context(), filter, &curencies); err != nil {
			return err
		}

		return c.JSON(200, curencies)
	})
}
