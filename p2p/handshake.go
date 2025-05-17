package p2p

// Handshake func is responsible for performing handshake between two nodes
type Handshakefunc func(Peer) error

func NOPHnaldshakeFunc(Peer) error {
	return nil
}
