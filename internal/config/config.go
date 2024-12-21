package config

import (
	"fmt"
	"os"

	"github.com/KarnerTh/xogs/internal/parser"
	"github.com/spf13/viper"
)

type Config struct {
	DefaultProfile string
	Profiles       []Profile
}

type Profile struct {
	Name          string
	Parser        string
	DisplayConfig DisplayConfig
}

func (c Config) GetProfileByName(name string) (*Profile, error) {
	for _, p := range c.Profiles {
		if p.Name == name {
			return &p, nil
		}
	}

	return nil, fmt.Errorf("Profile with name %s not found", name)
}

var DefaultProfile = Profile{
	Name:   "default",
	Parser: parser.ParserLogfmt, // TODO: sane default?
	DisplayConfig: DisplayConfig{
		Columns: []ColumnConfig{
			{
				Title:    "log",
				Width:    1,
				ValueKey: ValueKeyRaw,
			},
		},
	},
}

func Setup() Config {
	home, err := os.UserHomeDir()
	if err == nil {
		viper.AddConfigPath(home)
	}

	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".xogs")

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//TODO: default config?
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshalling config: %w", err))
	}
	return config
}
