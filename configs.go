// Copyright (c) 2020. Quimera Software S.p.A.

package sakura

import (
	"github.com/spf13/viper"
)

const (
	LanguageEng = iota
	LanguageEsp
)

type Config struct {
	AppName        string `mapstructure:"app_name"`
	LogPath        string `mapstructure:"log_path"`
	Language       int    `mapstructure:"_"` // Ignore it since we'll map it to an int
	BroadcastLevel int    `mapstructure:"broadcast_level"`
	LogLevel       int    `mapstructure:"log_level"`
	Discord        DiscordConfig
	Telegram       TelegramConfig
}

type DiscordConfig struct {
	Enabled   bool
	Username  string
	AvatarURL string `mapstructure:"avatar_url"`
	UseTTS    bool   `mapstructure:"use_tts"`
	Mentions  []int64
	Webhook   string
}

type TelegramConfig struct {
	Enabled  bool
	Token    string
	Channels []int64
}

func NewConfigFromFile(path string) (Config, error) {
	viper.SetDefault("log_path", "log.txt")
	viper.SetDefault("discord.username", "Sakura")
	viper.SetDefault("discord.avatar_url", "https://quimera.dev/sakura.jpg")
	viper.SetDefault("language", "eng")

	if path != "" {
		viper.SetConfigFile(path)
	} else {
		// Search in the working directory
		viper.SetConfigName("sakura")
		viper.AddConfigPath(".")
	}

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return Config{}, err
	}

	switch viper.GetString("language") {
	case "eng":
		cfg.Language = LanguageEng
	case "esp":
		cfg.Language = LanguageEsp
	}

	return cfg, nil
}
