package database

import (
	"sync"

	"github.com/Im-Stevemmmmm/bptree"
)

// LightDB is the database instance at runtime.
var LightDB *DB

// NewDB initializes a new DB from a slice of ReplicationNode.
func NewDB(replicationNodes []*ReplicationNode) *DB {
	return &DB{
		replicationNodes: replicationNodes,
		datafile:         NewRelativePath("ldb/data"),
		index:            bptree.NewTree(),
		ioWg:             &sync.WaitGroup{},
		ioRWMutex:        &sync.RWMutex{},
	}
}

// DB is the database configuration.
type DB struct {
	replicationNodes []*ReplicationNode
	datafile         *RelativePath
	index            *bptree.Tree
	ioWg             *sync.WaitGroup
	ioRWMutex        *sync.RWMutex
}
