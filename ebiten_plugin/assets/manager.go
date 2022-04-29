package assets

import "sync"

type Manager struct {
	mutex  *sync.Mutex
	assets map[string]Handle
}

func NewManager() *Manager {
	return &Manager{
		mutex:  &sync.Mutex{},
		assets: map[string]Handle{},
	}
}

func (m *Manager) Assets() map[string]Handle {
	return m.assets
}

func (m *Manager) Add(name string, asset Handle) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.assets[name] = asset
}

func (m *Manager) Get(name string) Handle {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.assets[name]
}

func (m *Manager) Delete(name string) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.assets, name)
}
