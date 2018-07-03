package main

import (
	"fmt"

	"lib/users"
	"lib/tokens"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type config struct {
	Server server
	Users  users.Config
	Tokens tokens.Config
}

type server struct {
	Port uint32
}

func (s server) addr() string {
	return fmt.Sprintf(":%d", s.Port)
}

func parseConfig() (*config, error) {
	v := viper.New()

	v.SetConfigName("gochat")

	v.AddConfigPath("/etc")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}

	c := &config{}
	err = v.UnmarshalExact(c)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse config to struct")
	}

	return c, nil
}
