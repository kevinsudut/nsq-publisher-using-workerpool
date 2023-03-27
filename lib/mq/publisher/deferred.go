package publisher

import (
	"time"

	jsoniter "github.com/json-iterator/go"
)

// DeferredPublish publishes raw byte to the given topic with specific delay time
func (mq *MessageQueue) DeferredPublish(topic string, delay time.Duration, body []byte) error {
	return mq.publishRetry(func() error {
		return mq.getClient().DeferredPublish(topic, delay, body)
	})
}

// DeferredPublish marshal the given data to JSON string using fast json encoder and then publish it to the given topic with specific delay time
func (mq *MessageQueue) DeferredPublishJSON(topic string, delay time.Duration, data interface{}) error {
	byt, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	return mq.Publish(topic, byt)
}

// DeferredPublish publishes raw string to the given topic with specific delay time
func (mq *MessageQueue) DeferredPublishString(topic string, delay time.Duration, data string) error {
	byt := []byte(data)

	return mq.Publish(topic, byt)
}
