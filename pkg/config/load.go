package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(r io.Reader) (Config, error) {
	var cfg Config
	err := yaml.NewDecoder(r).Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func LoadFile(filename string) (Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()
	return Load(f)
}
