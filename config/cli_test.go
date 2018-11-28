package config

import "testing"

func TestPortFromCli(t *testing.T) {
	t.Run("port is not empty", func(t *testing.T) {
		if GetPortFromCli() == "" {
			t.Error("Got empty string. Expected not empty")
		}
	})
}
