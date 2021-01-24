package kabusapi

import "golang.org/x/net/websocket"

func (c Client) SubscribePushStream() *PushStreamScanner {
	return NewPushStreamScanner(c.pushBaseURL.String())
}

func NewPushStreamScanner(url string) *PushStreamScanner {
	ws, err := websocket.Dial(url, "", "http://example.com")
	return &PushStreamScanner{
		ws:  ws,
		err: err,
	}
}

type PushStreamScanner struct {
	ws   *websocket.Conn
	err  error
	data BoardSuccess
}

func (s *PushStreamScanner) Data() BoardSuccess {
	return s.data
}

func (s *PushStreamScanner) Err() error {
	return s.err
}

func (s *PushStreamScanner) Scan() bool {
	if s.err != nil {
		return false
	}
	s.err = websocket.JSON.Receive(s.ws, &s.data)
	return s.err == nil
}

func (s *PushStreamScanner) Close() error {
	if s.ws == nil {
		return nil
	}
	return s.ws.Close()
}
