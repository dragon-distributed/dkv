package store



type PeerStore interface {
	// Peers returns the list of known peers.
	Peers() ([]string, error)

	// SetPeers sets the list of known peers. This is invoked when a peer is
	// added or removed.
	SetPeers([]string) error



}

