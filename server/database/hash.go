package database

import (
	"hash/fnv"
)

func hashOf(key string) int {
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32())
}
