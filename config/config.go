package config

type Config struct {
	Port string
}

func Get() (*Config, error) {
	return &Config{Port: "12345"}, nil
}
