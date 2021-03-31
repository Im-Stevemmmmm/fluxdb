package main

import (
	"github.com/Im-Stevemmmmm/bptree"

	"github.com/Im-Stevemmmmm/fluxdb/database"
)

const (
	defaultPort = 1623
)

func main() {
	db := &database.DB{
		ReplicationNodes: make([]*database.ReplicationNode, 0),
		KVPairs:          bptree.NewTree(),
	}

	if node := db.VerifyReplicationNodes(); node != nil {
		panic("invalid replication node: " + *node)
	}

	initAPI()
	displayStartupMessage()
}
