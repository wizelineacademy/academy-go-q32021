package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type server struct {
	Address string `json:"address"`
}

type config struct {
	Server server `json:"server"`
}

func ReadConfig(path string) (config, error) {

	jsonf, err := os.Open(path)

	if err != nil {
		return *new(config), err
	}

	data, err := ioutil.ReadAll(jsonf)

	if err != nil {
		return *new(config), err
	}

	var Config config

	err = json.Unmarshal(data, &Config)
	if err != nil {
		return *new(config), err
	}

	defer jsonf.Close()

	if (config{}) == Config {
		return config{}, errors.New("there was a problem reading the config file, please check.")
	}

	return Config, nil
}
