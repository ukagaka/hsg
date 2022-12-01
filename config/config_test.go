package config

import (
	"fmt"
	"testing"
)

func TestGetConfig(t *testing.T) {

	var cfg Configuration
	var err error
	cfg, err = Load("config.json")
	if err != nil {
		panic(err)
	}
	cfg.GetInt("level")

	fmt.Println("babaaaaaaaaaaaaaa", cfg)
}
