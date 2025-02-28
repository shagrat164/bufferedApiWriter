package buffer

import (
	"sync"

	"github.com/shagrat164/bufferedApiWriter/internal/model"
)

// Buffer - безопасная очередь.
type Buffer struct {
	data []model.Fact
	mu   sync.Mutex
}

func NewBuffer() *Buffer {
	return &Buffer{}
}

// Put записывает данные в очередь.
func (buf *Buffer) Put(record model.Fact) {
	buf.mu.Lock()
	defer buf.mu.Unlock()
	buf.data = append(buf.data, record)
}

// Get извлекает данные из очереди и true если данные есть, пустую структуру и false если буфер пуст.
func (buf *Buffer) Get() (model.Fact, bool) {
	buf.mu.Lock()
	defer buf.mu.Unlock()
	if len(buf.data) == 0 {
		return model.Fact{}, false
	}
	fact := buf.data[0]
	buf.data = buf.data[1:]
	return fact, true
}
