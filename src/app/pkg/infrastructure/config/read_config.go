package config

import (
	"github.com/pelletier/go-toml"
	"gonference/pkg/utils"
	"io/ioutil"
)

func ReadConfig() *Config {
	var config *Config
	data := readConfigFile()
	utils.Check(toml.Unmarshal(data, config))
	return config
}

func readConfigFile() []byte {
	data, err := ioutil.ReadFile("./config.toml"); utils.Check(err)
	return data
}
