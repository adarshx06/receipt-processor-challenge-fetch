package storage

import (
	"receipt-processor-challenge-fetch/models"
	"sync"
)

type MemoryStore struct {
	receipts map[string]models.Receipt
	points   map[string]int64
	mutex    sync.RWMutex
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		receipts: make(map[string]models.Receipt),
		points:   make(map[string]int64),
	}
}

func (s *MemoryStore) SaveReceipt(id string, receipt models.Receipt) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.receipts[id] = receipt
}

func (s *MemoryStore) SavePoints(id string, points int64) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.points[id] = points
}

func (s *MemoryStore) GetPoints(id string) (int64, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	points, exists := s.points[id]
	return points, exists
}
