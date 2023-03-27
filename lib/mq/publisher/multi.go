package publisher

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/kevinsudut/nsq-publisher-using-workerpool/util"
)

// MultiPublish publishes slice of raw byte to the given topic
func (mq *MessageQueue) MultiPublish(topic string, body [][]byte) error {
	return mq.publishRetry(func() error {
		return mq.getClient().MultiPublish(topic, body)
	})
}

// MultiPublishJSON marshal the given slice of data to JSON string using fast json encoder and then publish it to the given topic
func (mq *MessageQueue) MultiPublishJSON(topic string, data interface{}) error {
	btys := [][]byte{}

	res, err := util.ConvertInterfaceToSliceOfInterface(data)
	if err != nil {
		return err
	}

	for i := 0; i < len(res); i++ {
		byt, err := jsoniter.Marshal(res[i])
		if err != nil {
			return err
		}

		btys = append(btys, byt)
	}

	return mq.publishRetry(func() error {
		return mq.MultiPublish(topic, btys)
	})
}

// MultiPublish publishes slice of raw string to the given topic
func (mq *MessageQueue) MultiPublishString(topic string, data []string) error {
	btys := [][]byte{}

	for i := 0; i < len(data); i++ {
		btys = append(btys, []byte(data[i]))
	}

	return mq.publishRetry(func() error {
		return mq.MultiPublish(topic, btys)
	})
}
