package mq

import (
	"github.com/nsqio/go-nsq"
)

const (
	// Default NSQ port
	DefaultNSQPort = 4150

	// Default publish retry
	DefaultPubRetry = 3
)

// MessageQueueConfig is config for nsq producer
type MessageQueueConfig struct {
	Addresses []string `yaml:"addresses"`
	Port      int      `yaml:"port"`
	PubRetry  int      `yaml:"pub_retry"`

	NSQConfig *nsq.Config
}

// setDefaultConfig to set default value from message queue config
func (cfg *MessageQueueConfig) setDefaultConfig() {
	if cfg.NSQConfig == nil {
		cfg.NSQConfig = nsq.NewConfig()
	}

	if cfg.Port <= 0 {
		cfg.Port = DefaultNSQPort
	}

	if cfg.PubRetry <= 0 {
		cfg.PubRetry = DefaultPubRetry
	}
}
