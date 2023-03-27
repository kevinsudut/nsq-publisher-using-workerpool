package mq

import (
	"fmt"
	"log"
	"os"

	"github.com/nsqio/go-nsq"

	"github.com/kevinsudut/nsq-publisher-using-workerpool/lib/mq/publisher"
)

// newProducers creates new multiple nsq publisher
func newProducers(cfg MessageQueueConfig) (Publisher, error) {
	mq := publisher.InitPublisher(len(cfg.Addresses), cfg.PubRetry)

	for i := 0; i < len(cfg.Addresses); i++ {
		var err error

		producer, err := newProducer(cfg.Addresses[i], cfg.Port, cfg.NSQConfig)
		if err != nil {
			return nil, err
		}

		mq.SetProducer(producer)
	}

	return mq, nil
}

// newProducer creates new a single nsq publisher
func newProducer(address string, port int, nsqConfig *nsq.Config) (*nsq.Producer, error) {
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%d", address, port), nsqConfig)
	if err != nil {
		return nil, err
	}

	producer.SetLogger(log.New(os.Stderr, "", log.Flags()), nsq.LogLevelError)

	return producer, nil
}
