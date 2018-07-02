package main // import "gochat/cmd/server"

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type config struct {
	Server server
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
