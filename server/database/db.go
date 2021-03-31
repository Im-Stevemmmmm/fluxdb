package database

import (
	"github.com/Im-Stevemmmmm/bptree"
)

func NewDB(replicationNodes []*ReplicationNode) *DB {
	return &DB{
		replicationNodes: replicationNodes,
		kvPairs:          bptree.NewTree(),
	}
}

// DB is the database configuration.
type DB struct {
	replicationNodes []*ReplicationNode
	kvPairs          *bptree.Tree
}
