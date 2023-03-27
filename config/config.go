package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func ReadConfig(file string) (*Config, error) {
	var conf Config

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
