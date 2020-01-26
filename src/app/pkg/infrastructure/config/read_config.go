package config

import (
	"github.com/pelletier/go-toml"
	"gonference/pkg/utils"
	"io/ioutil"
)

var Config = ReadConfig()

func ReadConfig() *RootConfig {
	config := &RootConfig{}
	data := readConfigFile()
	utils.Check(toml.Unmarshal(data, config))
	return config
}

func readConfigFile() []byte {
	data, err := ioutil.ReadFile("./config.toml"); utils.Check(err)
	return data
}
