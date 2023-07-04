package repo_test

import (
	"context"
	"encoding/json"
	"log"

	"l0_wb/internal/adapters/repository"
	"l0_wb/internal/domain"
)

func main() {
	ctx := context.Background()
	dsn := "host=localhost port=5432 user=l0_back_user password=l0_back_hard_password dbname=l0_back_db sslmode=disable"
	repo, err := repository.New(ctx, dsn)
	if err != nil {
		log.Fatalf("can't create repository: %v", err)
	}

	models, err := repo.GetAll(ctx)
	if err != nil {
		log.Fatalf("can't get all models: %v", err)
	}

	log.Println("MODELS:", models)

	uid := "uid"
	model, err := repo.Get(ctx, uid)
	if err != nil {
		log.Fatalf("can't get model by uid = %s: %v", uid, err)
	}

	log.Println("MODEL:", model)

	var addModel domain.Model
	err = json.Unmarshal([]byte(testModel), &addModel)
	if err != nil {
		log.Fatalf("can't unmarshall model: %v", err)
	}

	err = repo.Add(ctx, addModel)
	if err != nil {
		log.Fatalf("can't add model: %v", err)
	}

	models, err = repo.GetAll(ctx)
	if err != nil {
		log.Fatalf("can't get all models: %v", err)
	}

	log.Println("MODELS:", models)
}

var testModel = `
{
  "order_uid": "b563feb7b2b84b6test123",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ],
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1"
}
`
