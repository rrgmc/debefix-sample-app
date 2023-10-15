package config

type StorageConfig struct {
	DatabaseURL string `yaml:"databaseURL"`
}

type Config struct {
	Storage StorageConfig `yaml:"storage"`
}
