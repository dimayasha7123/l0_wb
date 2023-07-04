package nats_subscriber

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/stan.go"
)

func (s *subscriber) BeginListen(ctx context.Context) error {
	subs, err := s.conn.Subscribe(
		s.subj,
		func(msg *stan.Msg) { s.msgHandlerWrapper(ctx, msg) },
		stan.SetManualAckMode(),
		stan.MaxInflight(1),
		stan.DurableName(s.durableName),
		stan.AckWait(3*time.Second),
	)
	if err != nil {
		s.conn.Close()
		return fmt.Errorf("can't subscribe to subj = %s: %v", s.subj, err)
	}

	s.subs = subs

	return nil
}
