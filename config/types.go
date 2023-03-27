package config

import (
	"github.com/kevinsudut/nsq-publisher-using-workerpool/lib/mq"
)

type Config struct {
	NSQ NSQConfig `yaml:"nsq"`
	CSV CSVConfig `yaml:"csv"`
	Log LogConfig `yaml:"log"`
}

type NSQConfig struct {
	MessageQueue           mq.MessageQueueConfig `yaml:"message_queue"`
	TopicName              string                `yaml:"topic_name"`
	NumPayloadMultiPublish int                   `yaml:"num_payload_multi_publish"`
	NumWorker              int                   `yaml:"num_worker"`
}

type CSVConfig struct {
	Path     string `yaml:"path"`
	FileName string `yaml:"file_name"`
}

type LogConfig struct {
	Path     string `yaml:"path"`
	LogLevel string `yaml:"log_level"`
}
