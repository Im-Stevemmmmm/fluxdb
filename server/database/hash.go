package database

import (
	"hash/fnv"
)

func hashOf(key string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return h.Sum32()
}
