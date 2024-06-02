package config

import "github.com/BurntSushi/toml"

type Config struct {
	Logger     LoggerConf
	GRPCServer GRPCConf
	DB         DBConf
	Rabbit     RabbitConf
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

type RabbitConf struct {
	URI      string
	Queue    string
	Exchange string
	Kind     string
}

func NewConfig(configFile string) (c Config, err error) {
	_, err = toml.DecodeFile(configFile, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
