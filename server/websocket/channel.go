package websocket

type WsChannel struct {
	ID         string `json:"id"`
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *WsMessage
}

func NewWsChannel(id string) *WsChannel {
	return &WsChannel{
		ID:         id,
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *WsMessage),
	}
}

func (channel *WsChannel) RunWsChannel() {
	for {
		select {
		case client := <-channel.register:
			channel.registerClientInWsChannel(client)

		case client := <-channel.unregister:
			channel.unregisterClientInWsChannel(client)

		case message := <-channel.broadcast:
			channel.broadcastToClientsInWsChannel(message.encode())
		}
	}
}

func (channel *WsChannel) registerClientInWsChannel(client *Client) {
	channel.clients[client] = true
}

func (channel *WsChannel) unregisterClientInWsChannel(client *Client) {
	delete(channel.clients, client)
}

func (channel *WsChannel) broadcastToClientsInWsChannel(message []byte) {
	for client := range channel.clients {
		client.send <- message
	}
}

func (channel *WsChannel) GetId() string {
	return channel.ID
}
