package kabusapi

import "golang.org/x/net/websocket"

type PushStreamReceiver struct {
	ws *websocket.Conn
}

func NewPushStreamReceiver(url string) (*PushStreamReceiver, error) {
	ws, err := websocket.Dial(url, "", "http://example.com")
	if err != nil {
		return nil, err
	}
	return &PushStreamReceiver{
		ws: ws,
	}, nil
}

func (c Client) SubscribePushStream() (*PushStreamReceiver, error) {
	return NewPushStreamReceiver(c.pushBaseURL.String())
}

func (s *PushStreamReceiver) ReceiveMessage() ([]byte, error) {
	var data []byte
	if err := websocket.Message.Receive(s.ws, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *PushStreamReceiver) Receive() (BoardSuccess, error) {
	var data BoardSuccess
	if err := websocket.JSON.Receive(s.ws, &data); err != nil {
		return BoardSuccess{}, err
	}
	return data, nil
}

func (s *PushStreamReceiver) Close() error {
	if s.ws == nil {
		return nil
	}
	return s.ws.Close()
}
