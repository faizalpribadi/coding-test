package config

import "github.com/spf13/viper"

var AppRootConfig *Config

type Config struct {
	App        AppConfig        `mapstructure:"app"`
	AIProvider AIProviderConfig `mapstructure:"ai_provider"`
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
}

type AIProviderConfig struct {
	Gemini AIGeminiConfig `mapstructure:"gemini"`
}

type AIGeminiConfig struct {
	ApiKey string `mapstructure:"api_key"`
	Model  string `mapstructure:"model"`
}

// Config holds the application configuration
func LoadConfig(file string) (*Config, error) {
	viper.SetConfigFile(file)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
