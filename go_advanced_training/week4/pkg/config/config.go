package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

func LoadConfigFromFile(path string, out interface{}) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read file error: %w", err)
	}
	err = yaml.Unmarshal(bytes, out)
	if err != nil {
		return fmt.Errorf("yaml unmarshal error: %w", err)
	}
	return nil
}
