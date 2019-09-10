package main_test

import (
	"testing"

	"github.com/spf13/viper"
)

func TestEnv(t *testing.T) {

	viper.SetConfigType("toml")
	viper.SetConfigFile(`config.toml`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

