package Studienarbeit_src

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	HandshakeTimeout: 0,
	ReadBufferSize:   2048,
	WriteBufferSize:  2048,
	WriteBufferPool:  nil,
	Subprotocols:     nil,
	Error:            nil,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	EnableCompression: false,
}

type WebSocket struct {
	Conn   *websocket.Conn
	Out    chan []byte
	In     chan []byte
	Events map[string]EventHandler
}

func NewWebSocket(w http.ResponseWriter, r *http.Request) (*WebSocket, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("An error accourde while upgrading the connection: %v", err)
		return nil, err
	}

	ws := &WebSocket{
		Conn:   conn,
		Out:    make(chan []byte),
		In:     make(chan []byte),
		Events: make(map[string]EventHandler),
	}

	go ws.Reader()
	go ws.Writer()
	return ws, nil
}

func (ws *WebSocket) Reader() {
	defer func() {
		ws.Conn.Close()
	}()
	for {
		_, message, err := ws.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("WS Message Error: %v", err)
			}
			break
		}
		event, err := NewEventFromRaw(message)
		if err != nil {
			log.Printf("Error parsing message: %v", err)
		} else {
			//
			log.Printf("MSG: %v", event)
		}
		if action, ok := ws.Events[event.Name]; ok {
			action(event)
		}
	}
}

func (ws *WebSocket) Writer() {
	for {
		select {
		case message, ok := <-ws.Out:
			if !ok {
				ws.Conn.WriteMessage(websocket.CloseMessage, make([]byte, 0))
				return
			}
			w, err := ws.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)
			w.Close()
		}
	}
}

func (ws *WebSocket) On(eventName string, action EventHandler) *WebSocket {
	ws.Events[eventName] = action
	return ws
}
