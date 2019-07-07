package cache

import (
	"github.com/YafimK/go-elastic-server-endpoint-server/model"
	"sync"
)

type DocumentCache struct {
	QueryCache     map[string]model.Documents
	DocumentsCache model.Documents
	lock           sync.RWMutex
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

func (dc DocumentCache) LookupQueryByFieldCache(query map[string]string) model.Documents {
	dc.lock.Lock()
	defer dc.lock.Unlock()

	//resultIndex := sort.Search(len(dc.DocumentsCache), func(i int) bool {
	//	if query["ip"] != "" && dc.DocumentsCache[i].Ip != ip {
	//		return false
	//	}
	//	if timestamp != "" && dc.DocumentsCache[i].Timestamp != timestamp {
	//		return false
	//	}
	//	if domain != "" && dc.DocumentsCache[i].Domain != domain {
	//		return false
	//	}
	//	if  != "" && dc.DocumentsCache[i].Domain != domain {
	//		return false
	//	}
	//	return false
	//})

	return nil
}
