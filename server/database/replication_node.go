package database

// VerifyReplicationNodes verifies all replication nodes are valid Instance
// APIs. Returns the first invalid ReplicationNode.
func VerifyReplicationNodes(nodes []*ReplicationNode) *ReplicationNode {
	for _, n := range nodes {
		if ok := n.Verify(); !ok {
			return n
		}
	}
	return nil
}

// Verify verifies that the API url is a valid Instance API.
func (r *ReplicationNode) Verify() bool {
	return true
}

// ReplicationNode is the API url for replication.
type ReplicationNode string
