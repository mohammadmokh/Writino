package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Redis RedisCfg
	Email EmailCfg
	Mongo MongoCfg
}

func LoadCfg(filename string) (Config, error) {

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return Config{}, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(buf, cfg)

	return *cfg, err
}
