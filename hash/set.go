package hash

import (
	"fmt"
	"sync"
)

type Check func(old interface{}, n interface{}) bool
type Hash struct {
	m map[string]interface{}
	sync.RWMutex
}

func (h *Hash) init() {
	if h.m == nil {
		h.m = make(map[string]interface{})
	}
}
func (h *Hash) Raw() map[string]interface{} {
	h.RWMutex.Lock()
	defer h.RWMutex.Unlock()
	h.init()
	r := make(map[string]interface{})
	for k, v := range h.m {
		r[k] = v
	}
	return r
}
func (h *Hash) Len() int {
	h.RWMutex.Lock()
	defer h.RWMutex.Unlock()
	h.init()
	return len(h.m)
}
func (h *Hash) Set(k string, i interface{}) {
	h.RWMutex.Lock()
	defer h.RWMutex.Unlock()
	h.init()
	h.m[k] = i
}
func (h *Hash) Get(k string) interface{} {
	h.RWMutex.RLock()
	defer h.RWMutex.RUnlock()
	h.init()
	if v, ok := h.m[k]; ok {
		return v
	} else {
		return nil
	}
}
func (h *Hash) Diff(hash *Hash, check ...Check) (lack, added *Hash) {
	raw := h.Raw()
	newHash := make(map[string]interface{})
	for k, v := range hash.Raw() {
		if v2, ok := raw[k]; ok {
			if len(check) >= 0 {
				for _, f := range check {
					if !f(v2, v) {
						goto B
					}
				}
			}
			delete(raw, k)
			continue
		}
	B:
		newHash[k] = v
	}
	lack = NewHash()
	lack.m = raw
	added = NewHash()
	added.m = newHash
	return
}
func (h Hash) String() string {
	h.RWMutex.RLock()
	defer h.RWMutex.RUnlock()
	return fmt.Sprint(h.m)
}
func NewHash() *Hash {
	return &Hash{
		m:       map[string]interface{}{},
		RWMutex: sync.RWMutex{},
	}
}
