package pkg

import "sync"

// Store 是一个线程安全的 map 封装，支持基本的增删改查操作。
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewStore 创建一个新的 Store 实例。
func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

// Set 设置 key 对应的值。
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get 获取 key 对应的值，若不存在返回空字符串和 false。
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}

// Delete 删除 key。
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

// List 返回所有的 key-value 对。
func (s *Store) List() map[string]string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	// 返回副本，避免外部修改
	result := make(map[string]string, len(s.data))
	for k, v := range s.data {
		result[k] = v
	}
	return result
}
