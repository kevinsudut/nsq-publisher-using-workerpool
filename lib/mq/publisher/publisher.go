package publisher

import (
	"sync"

	"github.com/nsqio/go-nsq"
)

// InitPublisher for initialize MessageQueue
func InitPublisher(producerCount int, pubRetry int) *MessageQueue {
	return &MessageQueue{
		producers:     make([]*nsq.Producer, producerCount),
		producerIdx:   0,
		producerCount: producerCount,
		mux:           sync.RWMutex{},
		pubRetry:      pubRetry,
	}
}

// SetProducer set producer entity to producers attributes
func (mq *MessageQueue) SetProducer(producer *nsq.Producer) {
	mq.producers[mq.producerIdx] = producer

	mq.setNextProducerIndex()
}

// getClient selects one of the available NSQ producers using round-robin load balancing.
func (mq *MessageQueue) getClient() *nsq.Producer {
	mq.mux.Lock()
	defer mq.mux.Unlock()

	producer := mq.producers[mq.producerIdx]

	mq.setNextProducerIndex()

	return producer
}

// setNextProducerIndex for set next producer index
func (mq *MessageQueue) setNextProducerIndex() {
	mq.producerIdx = (mq.producerIdx + 1) % mq.producerCount
}

// publishRetry for publish with retry mechanism
func (mq *MessageQueue) publishRetry(fn func() error) error {
	var err error

	for i := 0; i < mq.pubRetry; i++ {
		err = fn()
		if err == nil {
			break
		}
	}

	return err
}
