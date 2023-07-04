package nats_subscriber

func (s *subscriber) Close() {
	//s.subs.Unsubscribe()
	s.subs.Close()
	s.conn.Close()
}
