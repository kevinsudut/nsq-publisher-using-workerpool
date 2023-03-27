package mq

import (
	"errors"
	"time"
)

// Publisher is interface that must be implemented by nsq producer
type Publisher interface {
	// Publish
	Publish(topic string, body []byte) error
	PublishJSON(topic string, data interface{}) error
	PublishString(topic string, data string) error

	// Publish Multi
	MultiPublish(topic string, body [][]byte) error
	MultiPublishJSON(topic string, data interface{}) error
	MultiPublishString(topic string, data []string) error

	// Publish Defer
	DeferredPublish(topic string, delay time.Duration, body []byte) error
	DeferredPublishJSON(topic string, delay time.Duration, data interface{}) error
	DeferredPublishString(topic string, delay time.Duration, data string) error
}

// InitProducer for initialize nsq publisher
func InitProducer(cfg MessageQueueConfig) (Publisher, error) {
	if len(cfg.Addresses) == 0 {
		return nil, errors.New("addresses can not be empty")
	}

	cfg.setDefaultConfig()

	return newProducers(cfg)
}
