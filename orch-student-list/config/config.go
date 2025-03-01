package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config Struct
type Config struct {
	StudentAPIURL string
}

// LoadConfig อ่านค่าจาก ENV หรือไฟล์ config.yaml
func LoadConfig() *Config {
	viper.SetDefault("STUDENT_API_URL", "http://core-student-list.core.svc.cluster.local:8080")

	viper.AutomaticEnv()

	config := &Config{
		StudentAPIURL: viper.GetString("STUDENT_API_URL"),
	}

	log.Printf("Config Loaded: %+v\n", config)
	return config
}
