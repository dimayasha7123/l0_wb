package repository

import (
	"context"
	"fmt"

	"l0_wb/internal/domain"
)

func (r repository) Add(ctx context.Context, model domain.Model) error {
	query := `
	insert into models (order_uid, track_number, entry, delivery, payment, items, locale, internal_signature,
	                    customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14);
	`

	packed, err := packModel(model)
	if err != nil {
		return fmt.Errorf("can't pack model to query: %v", err)
	}

	_, err = r.pool.Exec(ctx, query, packed...)
	if err != nil {
		return fmt.Errorf("can't insert model: %v", err)
	}

	return nil
}
