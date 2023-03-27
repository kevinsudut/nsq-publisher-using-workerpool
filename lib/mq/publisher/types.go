package publisher

import (
	"sync"

	"github.com/nsqio/go-nsq"
)

// Message Queue defines nsq producer
type MessageQueue struct {
	producers     []*nsq.Producer
	producerIdx   int
	producerCount int
	mux           sync.RWMutex
	pubRetry      int
}
