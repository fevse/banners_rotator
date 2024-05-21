package config

import "github.com/BurntSushi/toml"

type Config struct {
	Logger     LoggerConf
	GRPCServer GRPCConf
	DB         DBConf
}

type LoggerConf struct {
	Level string
}

type GRPCConf struct {
	Network string
	Address string
}

type DBConf struct {
	Migration string
	DSN       string
}

func NewConfig(configFile string) (c Config, err error) {
	_, err = toml.DecodeFile(configFile, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
