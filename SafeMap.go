package SafeMap

import (
	"sync"
)

type BeeMap struct {
	lock *sync.RWMutex
	bm   map[interface{}]interface{}
}

func NewBeeMap() *BeeMap {
	return &BeeMap{
		lock: new(sync.RWMutex),
		bm:   make(map[interface{}]interface{}),
	}
}

//Get from maps return the k's value
func (m *BeeMap) Get(k interface{}) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if val, ok := m.bm[k]; ok {
		return val
	}
	return nil
}

//Get Size
func (m *BeeMap) Size() int {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return len(m.bm)
}

// Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *BeeMap) Set(k interface{}, v interface{}) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	if val, ok := m.bm[k]; !ok {
		m.bm[k] = v
	} else if val != v {
		m.bm[k] = v
	} else {
		return false
	}
	return true
}

// Returns true if k is exist in the map.
func (m *BeeMap) Check(k interface{}) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	if _, ok := m.bm[k]; !ok {
		return false
	}
	return true
}

func (m *BeeMap) Delete(k interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.bm, k)
}

func (m *BeeMap) Items() map[interface{}]interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.bm
}

//遍历每个元素
func (m *BeeMap) EachItem(eachFun func(interface{}, interface{})) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for key, value := range m.bm {
		eachFun(key, value)
	}
}

func (m *BeeMap) EachItemBreak(eachFun func(interface{}, interface{}) bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	for key, value := range m.bm {
		r := eachFun(key, value)
		if r == true {
			break
		}
	}
}

// Maps the given key and value. Returns false
// if the key is already in the map and changes nothing.
func (m *BeeMap) SetMap(bm *BeeMap) {
	m.lock.Lock()
	defer m.lock.Unlock()

	bm.EachItem(func(k interface{}, v interface{}) {
		m.bm[k] = v
	})
}

func (m *BeeMap) SetDatas(ks []interface{}, vs []interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	for index, k := range ks {
		v := vs[index]
		if val, ok := m.bm[k]; !ok {
			m.bm[k] = v
		} else if val != v {
			m.bm[k] = v
		}
	}
}

func (m *BeeMap) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()
	for k, _ := range m.bm {
		//可能不安全
		delete(m.bm, k)
	}
}
