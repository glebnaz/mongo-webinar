package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/glebnaz/mongo-webinar/internal/db"
	"github.com/glebnaz/mongo-webinar/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"net/http"
	"time"
)

const url = "https://www.cbr-xml-daily.ru/daily_json.js"

type Resp struct {
	Date         time.Time         `json:"Date"`
	PreviousDate time.Time         `json:"PreviousDate"`
	PreviousURL  string            `json:"PreviousURL"`
	Timestamp    time.Time         `json:"Timestamp"`
	Valute       map[string]Valute `json:"Valute"`
}

type Valute struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int     `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	store := db.NewStoreController(client, db.CurrencyDB, db.CurrencyCollection)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer func() {
		errClose := resp.Body.Close()
		if errClose != nil {
			panic(errClose)
		}
	}()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var respData Resp

	err = json.Unmarshal(data, &respData)
	if err != nil {
		panic(err)
	}

	for _, val := range respData.Valute {
		curr := model.Currency{
			NumCode:  val.NumCode,
			CharCode: val.CharCode,
			Nominal:  val.Nominal,
			Name:     val.Name,
			Value:    val.Value,
			Previous: val.Previous,
			Date:     time.Now().Unix(),
		}
		id, err := store.InsertOne(ctx, curr)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Curr: %v, id: %v\n", curr, id)
	}
}
