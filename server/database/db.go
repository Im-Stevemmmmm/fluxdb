package database

import (
	"sync"

	"github.com/Im-Stevemmmmm/bptree"
)

// Instance is the database instance at runtime.
var Instance *DB

// NewDB initializes a new DB from a slice of ReplicationNode.
func NewDB(replicationNodes []*ReplicationNode) *DB {
	return &DB{
		replicationNodes: replicationNodes,
		path:             NewRelativePath("ldb"),
		index:            bptree.NewTree(),
		ioWg:             &sync.WaitGroup{},
		ioRWMutex:        &sync.RWMutex{},
	}
}

// DB is the database configuration.
type DB struct {
	replicationNodes []*ReplicationNode
	path             *RelativePath
	index            *bptree.Tree
	ioWg             *sync.WaitGroup
	ioRWMutex        *sync.RWMutex
}
