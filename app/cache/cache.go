package cache

import (
	"github.com/YafimK/go-elastic-server-endpoint-server/model"
	"sync"
)

type DocumentCache struct {
	QueryCache map[string]model.Documents
	lock       sync.RWMutex
}

func NewCache() *DocumentCache {
	return &DocumentCache{QueryCache: make(map[string]model.Documents)}
}

func (dc DocumentCache) LookupQueryCache(value string) model.Documents {
	dc.lock.RLock()
	defer dc.lock.RUnlock()

	if result, isFound := dc.QueryCache[value]; isFound {
		return result
	}

	return nil
}

func (dc *DocumentCache) InsertQuery(key string, value model.Documents) {
	dc.lock.Lock()
	defer dc.lock.Unlock()

	dc.QueryCache[key] = value
}
