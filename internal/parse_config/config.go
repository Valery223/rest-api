package parseconfig

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-required:"true"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-required:"true"`
}

func MustLoadConfig() *Config {
	var configPath string
	flag.StringVar(&configPath, "config_path", "", "Path to the configuration file")
	flag.Parse()

	if configPath == "" {
		log.Fatal("config_path flag is required")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("configuration file does not exist: %s", configPath)
	}

	config := &Config{}
	if err := cleanenv.ReadConfig(configPath, config); err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	return config
}
