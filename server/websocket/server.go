package websocket

type WsServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	channels   map[*WsChannel]bool
}

func NewWebsocketServer() *WsServer {
	return &WsServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		channels:   make(map[*WsChannel]bool),
	}
}

func (server *WsServer) Run() {
	for {
		select {
		case client := <-server.register:
			server.registerClient(client)

		case client := <-server.unregister:
			server.unregisterClient(client)

		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

func (server *WsServer) registerClient(client *Client) {
	server.clients[client] = true
}

func (server *WsServer) unregisterClient(client *Client) {
	delete(server.clients, client)
}

func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

func (server *WsServer) findChannelByID(id string) *WsChannel {
	var foundWsChannel *WsChannel

	for channel := range server.channels {
		if channel.GetId() == id {
			foundWsChannel = channel
			break
		}
	}

	return foundWsChannel
}

func (server *WsServer) createWsChannel(id string) *WsChannel {
	channel := NewWsChannel(id)

	go channel.RunWsChannel()

	server.channels[channel] = true

	return channel
}
