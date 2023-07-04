package nats_subscriber

import (
	"context"
	"encoding/json"
	"fmt"

	"l0_wb/internal/domain"
	"l0_wb/internal/utils/logger"

	"github.com/nats-io/stan.go"
)

func (s *subscriber) msgHandlerWrapper(ctx context.Context, msg *stan.Msg) {
	err := s.msgHandler(ctx, msg)
	logMsg := "ADDED"
	if err != nil {
		logMsg = "ERROR " + err.Error()
	}

	logger.Log().Infow(
		"STREAM_MSG",
		"status", logMsg,
		"msg", msg,
	)
}

func (s *subscriber) msgHandler(ctx context.Context, msg *stan.Msg) error {
	err := s.app.Validate(ctx, msg.Data)
	if err != nil {
		s.ack(msg)
		return fmt.Errorf("can't validate msg data: %v", err)
	}

	var model domain.Model
	err = json.Unmarshal(msg.Data, &model)
	if err != nil {
		s.ack(msg)
		return fmt.Errorf("can't unmarshall msg data: %v", err)
	}

	exists, err := s.app.Exists(ctx, model.OrderUID)
	if err != nil {
		return fmt.Errorf("can't check existence of model with id = %s: %v", model.OrderUID, err)
	}

	if exists {
		s.ack(msg)
		return fmt.Errorf("model with id = %s already exists", model.OrderUID)
	}

	err = s.app.Add(ctx, model)
	if err != nil {
		return fmt.Errorf("can't add model with id = %s: %v", model.OrderUID, err)
	}
	s.ack(msg)

	return nil
}

func (s *subscriber) ack(msg *stan.Msg) {
	err := msg.Ack()
	if err != nil {
		logger.Log().Errorf("can't ack message with seq = %d", msg.Sequence)
		return
	}
}
