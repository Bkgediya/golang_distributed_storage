package p2p

// Peer is an interface that represent the remote node
type Peer interface {
}

// transport is anything that handles the p2p communication
// between the node in the networks, This can be of form
// TCP, UDP, WebSocket, etc.
type Transport interface {
	ListenAndAccept() error
}
