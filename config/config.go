/*
Example config.json
{
    "Port": ":8081"
}
*/

package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Port string
}

var err error
var config Config

func Get(filename string) (Config, error) {
	configFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Print("Unable to read config file, switching to flag mode")
		return Config{}, err
	}

	err = json.Unmarshal(configFile, &config)
	if err != nil {
		log.Print("Invalid JSON, expecting port from command line flag.")
		return Config{}, err
	}

	return config, nil

}
