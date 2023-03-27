package publisher

import (
	jsoniter "github.com/json-iterator/go"
)

// Publish publishes raw byte to the given topic
func (mq *MessageQueue) Publish(topic string, body []byte) error {
	return mq.publishRetry(func() error {
		return mq.getClient().Publish(topic, body)
	})
}

// PublishJSON marshal the given data to JSON string using fast json encoder and then publish it to the given topic
func (mq *MessageQueue) PublishJSON(topic string, data interface{}) error {
	byt, err := jsoniter.Marshal(data)
	if err != nil {
		return err
	}

	return mq.Publish(topic, byt)
}

// Publish publishes raw string to the given topic
func (mq *MessageQueue) PublishString(topic string, data string) error {
	byt := []byte(data)

	return mq.Publish(topic, byt)
}
