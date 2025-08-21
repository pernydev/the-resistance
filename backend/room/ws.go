package room

import (
	"sync"

	"github.com/gorilla/websocket"
)

type WSSender struct {
	conn   *websocket.Conn
	mutex  sync.Mutex
	closed bool
}

func NewWSSender(conn *websocket.Conn) *WSSender {
	return &WSSender{
		conn:   conn,
		closed: false,
	}
}

func (s *WSSender) SendMessage(messageType int, data []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.conn.WriteMessage(messageType, data)
}

func (s *WSSender) Close() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	err := s.conn.Close()
	if err != nil {
		return err
	}
	s.closed = true
	return nil
}
