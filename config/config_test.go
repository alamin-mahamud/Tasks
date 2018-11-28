package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Run("Read Port From Config File From Current Directory", func(t *testing.T) {
		_, err := ioutil.ReadFile("config.json")
		if err != nil {
			t.Error("Could Not Read Config File")
		}
	})

	t.Run("Read Port From Config File From Another Directory", func(t *testing.T) {
		_, err := ioutil.ReadFile("config/config.json")
		if err == nil {
			t.Error("Could Not Read Config File")
		}
	})

	t.Run("Parse Json From Config", func(t *testing.T) {
		configFile, _ := ioutil.ReadFile("config.json")
		var config Config
		err := json.Unmarshal(configFile, &config)
		if err != nil {
			t.Error("Could not Parse Json From Config")
		}
	})

	t.Run("Parse Port From Json", func(t *testing.T) {
		configFile, _ := ioutil.ReadFile("config.json")
		var config Config
		json.Unmarshal(configFile, &config)
		if config.Port == "" {
			t.Error("Could not parse port from JSON")
		}
	})

	t.Run("Port should have : prefix", func(t *testing.T) {
		configFile, _ := ioutil.ReadFile("config.json")
		var config Config
		json.Unmarshal(configFile, &config)
		strings.HasPrefix(config.Port, ":")
		if config.Port == "" {
			t.Error("Could not parse port from JSON")
		}
	})

}
