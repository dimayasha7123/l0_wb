package repository

import (
	"encoding/json"
	"errors"
	"fmt"

	"l0_wb/internal/domain"

	"github.com/jackc/pgx/v5"
)

func packModel(model domain.Model) ([]any, error) {
	errs := make([]error, 0, 3)
	jsonDelivery, err := json.Marshal(model.Delivery)
	if err != nil {
		errs = append(errs, fmt.Errorf("can't marshal delivery to json: %v", err))
	}

	jsonPayment, err := json.Marshal(model.Payment)
	if err != nil {
		errs = append(errs, fmt.Errorf("can't marshal payment to json: %v", err))
	}

	jsonItems, err := json.Marshal(model.Items)
	if err != nil {
		errs = append(errs, fmt.Errorf("can't marshal items to json: %v", err))
	}

	if len(errs) != 0 {
		return nil, fmt.Errorf("can't marshal fields: %v", errors.Join(errs...))
	}

	ret := []any{
		model.OrderUID,
		model.TrackNumber,
		model.Entry,
		jsonDelivery,
		jsonPayment,
		jsonItems,
		model.Locale,
		model.InternalSignature,
		model.CustomerID,
		model.DeliveryService,
		model.Shardkey,
		model.SmID,
		model.DateCreated,
		model.OofShard,
	}

	return ret, nil
}

func scanUnpackModel(row pgx.Row) (domain.Model, error) {
	ret := domain.Model{}
	var (
		jsonDelivery string
		jsonPayment  string
		jsonItems    string
	)

	err := row.Scan(
		&ret.OrderUID,
		&ret.TrackNumber,
		&ret.Entry,
		&jsonDelivery,
		&jsonPayment,
		&jsonItems,
		&ret.Locale,
		&ret.InternalSignature,
		&ret.CustomerID,
		&ret.DeliveryService,
		&ret.Shardkey,
		&ret.SmID,
		&ret.DateCreated,
		&ret.OofShard,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Model{}, fmt.Errorf("no model in result")
		}
		return domain.Model{}, fmt.Errorf("can't scan model: %v", err)
	}

	errs := make([]error, 0, 3)
	err = json.Unmarshal([]byte(jsonDelivery), &ret.Delivery)
	if err != nil {
		errs = append(errs, fmt.Errorf("can't unmarshall jsonDelivery: %v", err))
	}
	err = json.Unmarshal([]byte(jsonPayment), &ret.Payment)
	if err != nil {
		errs = append(errs, fmt.Errorf("can't unmarshall jsonPayment: %v", err))
	}
	err = json.Unmarshal([]byte(jsonItems), &ret.Items)
	if err != nil {
		errs = append(errs, fmt.Errorf("can't unmarshall jsonItems: %v", err))
	}
	if len(errs) != 0 {
		return domain.Model{}, fmt.Errorf("can't unmarshall some json fields: %v", errors.Join(errs...))
	}

	return ret, nil
}
